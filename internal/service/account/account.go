package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/DKA-Go-Microservices/Core-Account/generated/proto/account"
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
	account.UnimplementedAccountServer
}

var (
	db   = "dka_account"
	coll = "account"
)

// Create implementation
func (s *Server) Create(ctx context.Context, req *account.AccountModel) (*account.CreateResponse, error) {
	// Log asal permintaan untuk debugging
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("Request: %s -> %s", p.Addr, p.LocalAddr)
	}

	/**
	 * @Validation for Request Credential Data Is Required
	 */
	if req.Credential == "" {
		return &account.CreateResponse{
			Status: false,
			Code:   http.StatusBadRequest,
			Msg:    "Credential Is Required",
			Error:  errors.New("credential Is Required").Error(),
		}, nil
	}
	/**
	 * @Validation for Request Info Data Is Required
	 */
	if req.Info == "" {
		return &account.CreateResponse{
			Status: false,
			Code:   http.StatusBadRequest,
			Msg:    "Info Is Required",
			Error:  errors.New("info Is Required").Error(),
		}, nil
	}

	// Koneksi ke database
	db, err := database.Client(ctx).GetDatabase(db)
	if err != nil {
		log.Println("Database connection error:", err)
		return &account.CreateResponse{
			Status: false,
			Code:   http.StatusInternalServerError,
			Msg:    "Failed to connect to database",
			Error:  err.Error(),
		}, nil
	}

	// Insert data ke koleksi
	res, err := db.Collection(coll).InsertOne(ctx, req)
	if err != nil {
		log.Println("Insert operation error:", err)
		return &account.CreateResponse{
			Status: false,
			Code:   http.StatusUnprocessableEntity,
			Msg:    "Failed to insert data into database",
			Error:  err.Error(),
		}, nil
	}

	return &account.CreateResponse{
		Status: true,
		Code:   http.StatusOK,
		Msg:    "Data inserted successfully",
		Id:     fmt.Sprintf("%v", res.InsertedID),
	}, nil
}

// Read Implementation
func (s *Server) Read(ctx context.Context, req *account.ReadRequest) (*account.ReadResponse, error) {
	startTime := time.Now() // Catat waktu mulai
	// Log the request origin for debugging
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("Request: %s -> %s", p.Addr, p.LocalAddr)
	}
	// Connect to database
	db, err := database.Client(ctx).GetDatabase(db)
	if err != nil {
		log.Println("Database connection error:", err)
		return &account.ReadResponse{
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
		return &account.ReadResponse{
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

	var results []*account.AccountModel
	// Use cursor.All to decode all documents at once
	if err := cursor.All(ctx, &results); err != nil { // Perhatikan penggunaan '&' di sini
		log.Println("Error decoding cursor:", err)
		return &account.ReadResponse{
			Status: false,
			Code:   http.StatusUnprocessableEntity,
			Msg:    "Failed to decode data",
			Error:  err.Error(),
		}, nil
	}

	// Hitung durasi waktu proses
	duration := time.Since(startTime)

	response := &account.ReadResponse{
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
	account.RegisterAccountServer(grpcServer, server)
}
