package main

import (
	"github.com/DKA-Go-Microservices/Core-Account/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Set Default Value If Env Not Defined
	if value := os.Getenv("DKA_SERVICE_PORT"); value == "" {
		_ = os.Setenv("DKA_SERVICE_PORT", "8080")
	}
	lis, err := net.Listen("tcp", "0.0.0.0:"+os.Getenv("DKA_SERVICE_PORT"))
	// Check Error Listen TCP Listen
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create New Server GPRC
	grpcServer := grpc.NewServer()
	// Add Goroutine Service
	go func() {
		// Adding Registration Services
		service.Service(grpcServer)
	}()
	// Channel to capture termination signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	// Goroutine for serving gRPC
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Println("SERVER: Successfully started microservices on " + os.Getenv("DKA_SERVICE_PORT"))
	// Wait for a termination signal
	<-stop
	// Gracefully stop the gRPC server
	log.Println("SERVER: Shutting down gracefully...")
	grpcServer.GracefulStop()
	// Optionally, wait a little to ensure shutdown is clean
	time.Sleep(time.Second)
	log.Println("SERVER: Stopped.")
}
