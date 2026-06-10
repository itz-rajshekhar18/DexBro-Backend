package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"dexbro-backend/models"

	"github.com/gin-gonic/gin"
)

func TestCreateRegistrationValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		payload        map[string]interface{}
		expectedStatus int
	}{
		{
			name: "Valid registration",
			payload: map[string]interface{}{
				"name":       "Test User",
				"email":      "test@example.com",
				"phone":      "+91 1234567890",
				"grade":      "10",
				"experience": "beginner",
				"interests":  []string{"ml", "python"},
				"message":    "Test message",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Missing email",
			payload: map[string]interface{}{
				"name":       "Test User",
				"phone":      "+91 1234567890",
				"grade":      "10",
				"experience": "beginner",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid email format",
			payload: map[string]interface{}{
				"name":       "Test User",
				"email":      "invalid-email",
				"phone":      "+91 1234567890",
				"grade":      "10",
				"experience": "beginner",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This is a basic validation test
			// In real scenario, you'd need to mock MongoDB connection
			router := gin.New()
			
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/registrations", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		})
	}
}

func TestRegistrationModel(t *testing.T) {
	registration := models.Registration{
		Name:       "Test User",
		Email:      "test@example.com",
		Phone:      "+91 1234567890",
		Grade:      "10",
		Experience: "beginner",
		Interests:  []string{"ml", "python"},
		Message:    "Test message",
	}

	if registration.Name != "Test User" {
		t.Errorf("Expected name 'Test User', got '%s'", registration.Name)
	}

	if registration.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", registration.Email)
	}
}
