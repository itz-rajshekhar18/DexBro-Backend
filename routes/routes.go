package routes

import (
	"net/http"

	"dexbro-backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "DexBro Workshop API is running",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Registration routes
		registrations := v1.Group("/registrations")
		{
			registrations.POST("", controllers.CreateRegistration)
			registrations.GET("", controllers.GetAllRegistrations)
			registrations.GET("/:id", controllers.GetRegistrationByID)
			registrations.DELETE("/:id", controllers.DeleteRegistration)
		}
	}

	// 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Route not found",
		})
	})
}
