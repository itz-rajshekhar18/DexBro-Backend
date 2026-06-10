package controllers

import (
	"context"
	"net/http"
	"time"

	"dexbro-backend/config"
	"dexbro-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const registrationCollection = "registrations"

// CreateRegistration handles POST /api/registrations
func CreateRegistration(c *gin.Context) {
	var registration models.Registration

	// Bind and validate JSON
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, models.RegistrationResponse{
			Success: false,
			Message: "Invalid request data: " + err.Error(),
		})
		return
	}

	// Set timestamps
	registration.CreatedAt = time.Now()
	registration.UpdatedAt = time.Now()

	// Insert into database
	collection := config.GetCollection(registrationCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, registration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.RegistrationResponse{
			Success: false,
			Message: "Failed to create registration: " + err.Error(),
		})
		return
	}

	registration.ID = result.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, models.RegistrationResponse{
		Success: true,
		Message: "Registration created successfully",
		Data:    &registration,
	})
}

// GetAllRegistrations handles GET /api/registrations
func GetAllRegistrations(c *gin.Context) {
	collection := config.GetCollection(registrationCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get total count
	total, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.RegistrationListResponse{
			Success: false,
			Message: "Failed to count registrations: " + err.Error(),
		})
		return
	}

	// Find all registrations, sorted by created_at descending
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.RegistrationListResponse{
			Success: false,
			Message: "Failed to fetch registrations: " + err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)

	var registrations []models.Registration
	if err := cursor.All(ctx, &registrations); err != nil {
		c.JSON(http.StatusInternalServerError, models.RegistrationListResponse{
			Success: false,
			Message: "Failed to decode registrations: " + err.Error(),
		})
		return
	}

	// Return empty array instead of null
	if registrations == nil {
		registrations = []models.Registration{}
	}

	c.JSON(http.StatusOK, models.RegistrationListResponse{
		Success: true,
		Message: "Registrations fetched successfully",
		Data:    registrations,
		Total:   total,
	})
}

// GetRegistrationByID handles GET /api/registrations/:id
func GetRegistrationByID(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.RegistrationResponse{
			Success: false,
			Message: "Invalid registration ID",
		})
		return
	}

	collection := config.GetCollection(registrationCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var registration models.Registration
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&registration)
	if err != nil {
		c.JSON(http.StatusNotFound, models.RegistrationResponse{
			Success: false,
			Message: "Registration not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.RegistrationResponse{
		Success: true,
		Message: "Registration fetched successfully",
		Data:    &registration,
	})
}

// DeleteRegistration handles DELETE /api/registrations/:id
func DeleteRegistration(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.RegistrationResponse{
			Success: false,
			Message: "Invalid registration ID",
		})
		return
	}

	collection := config.GetCollection(registrationCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.RegistrationResponse{
			Success: false,
			Message: "Failed to delete registration: " + err.Error(),
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, models.RegistrationResponse{
			Success: false,
			Message: "Registration not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.RegistrationResponse{
		Success: true,
		Message: "Registration deleted successfully",
	})
}
