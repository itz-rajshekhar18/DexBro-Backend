package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Registration represents a workshop registration
type Registration struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Email           string             `json:"email" bson:"email"`
	Phone           string             `json:"phone" bson:"phone"`
	Grade           string             `json:"grade" bson:"grade"`
	Experience      string             `json:"experience" bson:"experience"`
	Interests       []string           `json:"interests" bson:"interests"`
	Message         string             `json:"message" bson:"message"`
	PaymentStatus   string             `json:"payment_status" bson:"payment_status"` // pending, success, failed
	PaymentID       string             `json:"payment_id" bson:"payment_id"`
	OrderID         string             `json:"order_id" bson:"order_id"`
	RazorpayOrderID string             `json:"razorpay_order_id" bson:"razorpay_order_id"`
	Amount          int                `json:"amount" bson:"amount"` // Amount in paise (75000 paise = ₹750)
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}

// RegistrationResponse is the response format
type RegistrationResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *Registration `json:"data,omitempty"`
}

// RegistrationListResponse for listing registrations
type RegistrationListResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    []Registration  `json:"data"`
	Total   int64           `json:"total"`
}

// PaymentOrder represents a Razorpay order
type PaymentOrder struct {
	OrderID         string `json:"order_id"`
	RazorpayOrderID string `json:"razorpay_order_id"`
	Amount          int    `json:"amount"`
	Currency        string `json:"currency"`
	KeyID           string `json:"key_id"`
}

// PaymentOrderResponse for payment order creation
type PaymentOrderResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *PaymentOrder `json:"data,omitempty"`
}

// PaymentVerification represents payment verification request
type PaymentVerification struct {
	OrderID           string `json:"order_id" binding:"required"`
	RazorpayOrderID   string `json:"razorpay_order_id" binding:"required"`
	RazorpayPaymentID string `json:"razorpay_payment_id" binding:"required"`
	RazorpaySignature string `json:"razorpay_signature" binding:"required"`
}

// PaymentVerificationResponse for payment verification
type PaymentVerificationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

