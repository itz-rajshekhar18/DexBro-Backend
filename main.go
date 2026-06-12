package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dexbro-backend/config"
	"dexbro-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := config.ConnectDB(ctx); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer config.DisconnectDB(context.Background())

	log.Println("Successfully connected to MongoDB")

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize router
	router := gin.Default()

	// CORS configuration
	// Build allowed origins list
	allowedOrigins := []string{
		"http://localhost:3000",
		"https://dexbro-workshop.vercel.app",
		"https://dex-guru-workshop.vercel.app",
		"https://dex-bro-workshop.vercel.app",
	}
	
	// Add URLs from environment variables if they exist
	if frontendURL := getEnv("FRONTEND_URL", ""); frontendURL != "" {
		allowedOrigins = append(allowedOrigins, frontendURL)
	}
	if dexbroURL := getEnv("FRONTEND_DEXBRO_URL", ""); dexbroURL != "" {
		allowedOrigins = append(allowedOrigins, dexbroURL)
	}

	corsConfig := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))

	// Setup routes
	routes.SetupRoutes(router)

	// Get port from environment
	port := getEnv("PORT", "8080")

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := router.Run(":" + port); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
