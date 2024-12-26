package service

import (
	"github.com/DKA-Go-Microservices/Core-Account/internal/service/account"
	"google.golang.org/grpc"
)

func Service(server *grpc.Server) {
	// Initialize the service server
	srv := account.NewServer()
	// Register the service with the gRPC server
	account.RegisterGRPCServer(server, srv)
}
