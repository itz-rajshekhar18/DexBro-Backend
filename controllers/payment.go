package controllers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"dexbro-backend/config"
	"dexbro-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	WORKSHOP_AMOUNT = 75000 // ₹750 in paise (Razorpay uses paise)
	CURRENCY        = "INR"
)

// RazorpayOrderRequest represents Razorpay order creation request
type RazorpayOrderRequest struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`
}

// RazorpayOrderResponse represents Razorpay order creation response
type RazorpayOrderResponse struct {
	ID       string `json:"id"`
	Entity   string `json:"entity"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Receipt  string `json:"receipt"`
	Status   string `json:"status"`
}

// CreatePaymentOrder creates a Razorpay order
// POST /api/v1/payment/create-order
func CreatePaymentOrder(c *gin.Context) {
	var registration models.Registration
	
	// Bind JSON request
	if err := c.ShouldBindJSON(&registration); err != nil {
		log.Printf("Validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request data. Please check all required fields.",
			"error":   err.Error(),
			"details": "Required: name, email, phone, grade, experience",
		})
		return
	}

	// Validate required fields explicitly
	if registration.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Name is required",
		})
		return
	}
	if registration.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Email is required",
		})
		return
	}
	if registration.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Phone is required",
		})
		return
	}
	if registration.Grade == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Grade is required",
		})
		return
	}
	if registration.Experience == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Experience level is required",
		})
		return
	}

	// Initialize empty array for interests if nil
	if registration.Interests == nil {
		registration.Interests = []string{}
	}

	log.Printf("Creating payment order for: %s (%s)", registration.Name, registration.Email)

	// Generate unique order ID
	orderID := primitive.NewObjectID().Hex()
	
	// Create Razorpay order
	razorpayOrderID, err := createRazorpayOrder(orderID, WORKSHOP_AMOUNT)
	if err != nil {
		log.Printf("Failed to create Razorpay order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create payment order",
			"error":   err.Error(),
		})
		return
	}

	// Set payment fields
	registration.OrderID = orderID
	registration.RazorpayOrderID = razorpayOrderID
	registration.Amount = WORKSHOP_AMOUNT
	registration.PaymentStatus = "pending"
	registration.CreatedAt = time.Now()
	registration.UpdatedAt = time.Now()

	// Insert into database
	collection := config.GetCollection("registrations")
	result, err := collection.InsertOne(c.Request.Context(), registration)
	if err != nil {
		log.Printf("Failed to save registration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to save registration",
			"error":   err.Error(),
		})
		return
	}

	registration.ID = result.InsertedID.(primitive.ObjectID)

	log.Printf("Payment order created successfully: %s", razorpayOrderID)

	// Return payment order details
	c.JSON(http.StatusOK, models.PaymentOrderResponse{
		Success: true,
		Message: "Payment order created successfully",
		Data: &models.PaymentOrder{
			OrderID:         orderID,
			RazorpayOrderID: razorpayOrderID,
			Amount:          WORKSHOP_AMOUNT,
			Currency:        CURRENCY,
			KeyID:           os.Getenv("RAZORPAY_KEY_ID"),
		},
	})
}

// VerifyPayment verifies the Razorpay payment signature
// POST /api/v1/payment/verify
func VerifyPayment(c *gin.Context) {
	var verification models.PaymentVerification

	// Bind JSON request
	if err := c.ShouldBindJSON(&verification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	// Verify signature
	if !verifyPaymentSignature(verification.RazorpayOrderID, verification.RazorpayPaymentID, verification.RazorpaySignature) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid payment signature",
		})
		return
	}

	// Update payment status in database
	collection := config.GetCollection("registrations")
	filter := bson.M{"order_id": verification.OrderID}
	update := bson.M{
		"$set": bson.M{
			"payment_status": "success",
			"payment_id":     verification.RazorpayPaymentID,
			"updated_at":     time.Now(),
		},
	}

	result, err := collection.UpdateOne(c.Request.Context(), filter, update)
	if err != nil {
		log.Printf("Failed to update payment status: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update payment status",
			"error":   err.Error(),
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Order not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.PaymentVerificationResponse{
		Success: true,
		Message: "Payment verified successfully",
	})
}

// GetPaymentStatus gets the payment status for an order
// GET /api/v1/payment/status/:orderId
func GetPaymentStatus(c *gin.Context) {
	orderID := c.Param("orderId")

	collection := config.GetCollection("registrations")
	var registration models.Registration

	err := collection.FindOne(c.Request.Context(), bson.M{"order_id": orderID}).Decode(&registration)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Order not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"order_id":       registration.OrderID,
			"payment_status": registration.PaymentStatus,
			"amount":         registration.Amount,
		},
	})
}

// createRazorpayOrder creates an order in Razorpay
func createRazorpayOrder(orderID string, amount int) (string, error) {
	keyID := os.Getenv("RAZORPAY_KEY_ID")
	keySecret := os.Getenv("RAZORPAY_KEY_SECRET")

	if keyID == "" || keySecret == "" {
		return "", fmt.Errorf("Razorpay credentials not configured")
	}

	// Create order request
	orderRequest := RazorpayOrderRequest{
		Amount:   amount,
		Currency: CURRENCY,
		Receipt:  orderID,
	}

	jsonData, err := json.Marshal(orderRequest)
	if err != nil {
		return "", err
	}

	// Make API call to Razorpay
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/orders", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(keyID, keySecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call Razorpay API: %v", err)
	}
	defer resp.Body.Close()

	// Read response body for debugging
	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Printf("Razorpay API response status: %d, body: %s", resp.StatusCode, string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("razorpay API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse response
	var orderResponse RazorpayOrderResponse
	if err := json.Unmarshal(bodyBytes, &orderResponse); err != nil {
		return "", fmt.Errorf("failed to parse Razorpay response: %v", err)
	}

	return orderResponse.ID, nil
}

// verifyPaymentSignature verifies the Razorpay payment signature
func verifyPaymentSignature(orderID, paymentID, signature string) bool {
	keySecret := os.Getenv("RAZORPAY_KEY_SECRET")
	
	// Create signature string
	message := orderID + "|" + paymentID
	
	// Create HMAC SHA256
	h := hmac.New(sha256.New, []byte(keySecret))
	h.Write([]byte(message))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}


// DebugPaymentRequest helps debug what data is being received
// POST /api/v1/payment/debug
func DebugPaymentRequest(c *gin.Context) {
	// Get raw body
	bodyBytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to read request body",
			"error":   err.Error(),
		})
		return
	}

	// Try to parse as JSON
	var data map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid JSON",
			"error":   err.Error(),
			"rawBody": string(bodyBytes),
		})
		return
	}

	// Try to bind to registration model
	var registration models.Registration
	if err := json.Unmarshal(bodyBytes, &registration); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Debug info",
			"receivedData": data,
			"parseError":  err.Error(),
		})
		return
	}

	// Show what was received and parsed
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      "Debug info",
		"receivedData": data,
		"parsedAs": gin.H{
			"name":       registration.Name,
			"email":      registration.Email,
			"phone":      registration.Phone,
			"grade":      registration.Grade,
			"experience": registration.Experience,
			"interests":  registration.Interests,
			"message":    registration.Message,
		},
		"validation": gin.H{
			"name_valid":       registration.Name != "",
			"email_valid":      registration.Email != "",
			"phone_valid":      registration.Phone != "",
			"grade_valid":      registration.Grade != "",
			"experience_valid": registration.Experience != "",
		},
	})
}
