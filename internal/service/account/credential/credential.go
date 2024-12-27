package credential

import (
	"context"
	"github.com/DKA-Go-Microservices/Core-Account/generated/proto/account/credential"
	database "github.com/DKA-Go-Microservices/Core-Account/internal/database/MongoDB"
	"github.com/DKA-Go-Microservices/Core-Account/internal/helper"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"net/http"
	"time"
)

// Server struct implements the gRPC service
type Server struct {
	credential.UnimplementedCredentialServer
}

var (
	db   = "dka_account"
	coll = "account_credential"
)

// Read Implementation
func (s *Server) Read(ctx context.Context, req *credential.ReadRequest) (*credential.ReadResponse, error) {
	startTime := time.Now() // Catat waktu mulai
	// Log the request origin for debugging
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("Request: %s -> %s", p.Addr, p.LocalAddr)
	}

	// Connect to database client
	db, err := database.Client(ctx).GetDatabase(db)
	if err != nil {
		log.Println("Database connection error:", err)
		return &credential.ReadResponse{
			Status: false,
			Code:   http.StatusUnprocessableEntity,
			Msg:    "Failed To Get Data",
			Error:  err.Error(),
		}, nil
	}

	// Find documents in the collection
	cursor, err := db.Collection(coll).Find(ctx, bson.D{})
	if err != nil {
		log.Println("Find query error:", err)
		return &credential.ReadResponse{
			Status: false,
			Code:   http.StatusUnprocessableEntity,
			Msg:    "Failed to fetch data from database",
			Error:  err.Error(),
		}, nil
	}

	// Ensure cursor is closed when function returns
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Println("Error closing cursor:", err)
		}
	}()

	var results []*credential.CredentialModel
	// Use cursor.All to decode all documents at once
	if err := cursor.All(ctx, &results); err != nil { // Perhatikan penggunaan '&' di sini
		log.Println("Error decoding cursor:", err)
		return &credential.ReadResponse{
			Status: false,
			Code:   http.StatusUnprocessableEntity,
			Msg:    "Failed to decode data",
			Error:  err.Error(),
		}, nil
	}

	// Hitung durasi waktu proses
	duration := time.Since(startTime)

	response := &credential.ReadResponse{
		Status: true,
		Code:   http.StatusOK,
		Msg:    "Data fetched successfully",
		Data:   results,
	}

	log.Printf("Response: %s", helper.FormatDurationID(duration))

	return response, nil
}

// NewServer creates and initializes a new gRPC server
func NewServer() *Server {
	return &Server{}
}

// RegisterGRPCServer registers the gRPC service with the server
func RegisterGRPCServer(grpcServer *grpc.Server, server *Server) {
	credential.RegisterCredentialServer(grpcServer, server)
}
