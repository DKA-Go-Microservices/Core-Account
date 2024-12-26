package database

import (
	"context"
	"github.com/DKA-Go-Microservices/Core-Account/internal/connection/MongoDB"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type DB struct {
	ctx context.Context
}

func (db DB) GetDatabase(dbName string) (*mongo.Database, error) {
	// Add Default error nil
	var err error = nil
	if value := os.Getenv("DKA_DB_HOST"); value == "" {
		_ = os.Setenv("DKA_DB_HOST", "127.0.0.1")
	}
	if value := os.Getenv("DKA_DB_PORT"); value == "" {
		_ = os.Setenv("DKA_DB_PORT", "27017")
	}
	if value := os.Getenv("DKA_DB_NAME"); value == "" {
		_ = os.Setenv("DKA_DB_NAME", "dka_parking")
	}

	if value := os.Getenv("DKA_DB_USERNAME"); value == "" {
		_ = os.Setenv("DKA_DB_USERNAME", "root")
	}
	if value := os.Getenv("DKA_DB_PASSWORD"); value == "" {
		_ = os.Setenv("DKA_DB_PASSWORD", "")
	}

	// MongoDB configuration
	config := MongoDB.Config{
		URI:      "mongodb://" + os.Getenv("DKA_DB_HOST") + ":" + os.Getenv("DKA_DB_PORT"),
		Database: os.Getenv("DKA_DB_NAME"),
		Auth: &MongoDB.ConfigAuth{
			Username: os.Getenv("DKA_DB_USERNAME"),
			Password: os.Getenv("DKA_DB_PASSWORD"),
		},
	}
	// Get MongoDB client
	client, err := MongoDB.MongoDB(config)
	// Ensure the client disconnects on function exit
	dbInstance := client.Database(dbName)
	// Return the database instance and nil error
	return dbInstance, err
}

func Client(ctx context.Context) DB {
	return DB{ctx: ctx}
}
