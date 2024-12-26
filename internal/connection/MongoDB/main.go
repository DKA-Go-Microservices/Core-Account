package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

type ConfigAuth struct {
	Username string
	Password string
}

// Config MongoDBConfig contains MongoDB settings
type Config struct {
	URI      string
	Database string
	Port     int
	Auth     *ConfigAuth
}

var (
	mongoClient *mongo.Client = nil
	once        sync.Once     // Ensures that the MongoDB client is initialized only once
)

// MongoDB connects with the given configuration and ensures it's done only once
func MongoDB(config Config) (*mongo.Client, error) {
	var err error = nil
	// Use sync.Once to ensure MongoDB connection is created only once
	once.Do(func() {
		// Set up client options and authentication
		clientOptions := options.Client()
		clientOptions.ApplyURI(config.URI)
		clientOptions.SetTimeout(5 * time.Second)
		clientOptions.SetConnectTimeout(5 * time.Second)

		if config.Auth != nil {
			clientOptions.SetAuth(options.Credential{
				Username: config.Auth.Username,
				Password: config.Auth.Password,
			})
		}
		// Set timeout (optional)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Connect to MongoDB
		mongoClient, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("MongoDB: Connected")
		}

		// Check connection
		if err = mongoClient.Ping(ctx, nil); err != nil {
			log.Printf("mongo server is live ~")
		} else {
			log.Println("MongoDB: Ping test successfully")
		}

	})

	// Return the client and any error encountered
	return mongoClient, err
}
