package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

// ConnectDB establishes connection to MongoDB
func ConnectDB(ctx context.Context) error {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return fmt.Errorf("MONGODB_URI environment variable is not set")
	}

	dbName := os.Getenv("MONGODB_DATABASE")
	if dbName == "" {
		return fmt.Errorf("MONGODB_DATABASE environment variable is not set")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)
	
	// Configure TLS for MongoDB Atlas - bypass verification for Windows compatibility
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	clientOptions.SetTLSConfig(tlsConfig)
	
	// Set timeouts and pool size
	clientOptions.SetServerSelectionTimeout(10 * time.Second)
	clientOptions.SetConnectTimeout(10 * time.Second)
	clientOptions.SetMaxPoolSize(50)
	clientOptions.SetMinPoolSize(5)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	Client = client
	Database = client.Database(dbName)

	return nil
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB(ctx context.Context) error {
	if Client == nil {
		return nil
	}
	return Client.Disconnect(ctx)
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return Database.Collection(collectionName)
}
