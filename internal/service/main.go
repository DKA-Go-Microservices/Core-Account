package service

import (
	"github.com/DKA-Go-Microservices/Core-Account/internal/service/account"
	"github.com/DKA-Go-Microservices/Core-Account/internal/service/account/credential"
	"github.com/DKA-Go-Microservices/Core-Account/internal/service/account/info"
	"google.golang.org/grpc"
)

func Service(server *grpc.Server) {
	// Register the service with the gRPC server
	account.RegisterGRPCServer(server, account.NewServer())
	// Register the service with the gRPC server
	credential.RegisterGRPCServer(server, credential.NewServer())
	// Register Info
	info.RegisterGRPCServer(server, info.NewServer())
}
