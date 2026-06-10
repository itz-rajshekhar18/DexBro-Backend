package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Registration represents a workshop registration
type Registration struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Email      string             `json:"email" bson:"email" binding:"required,email"`
	Phone      string             `json:"phone" bson:"phone" binding:"required"`
	Grade      string             `json:"grade" bson:"grade" binding:"required"`
	Experience string             `json:"experience" bson:"experience" binding:"required"`
	Interests  []string           `json:"interests" bson:"interests"`
	Message    string             `json:"message" bson:"message"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
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
