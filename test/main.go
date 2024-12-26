package main

import (
	"context"
	"fmt"
	"github.com/DKA-Go-Microservices/Core-Account/generated/proto/account"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Connect to the gRPC server on port 5051
	conn, err := grpc.Dial("127.0.0.1:5052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	// Create a new ExampleService client
	client := account.NewAccountClient(conn)

	res, err := client.Read(context.Background(), &account.ReadRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print the response message
	fmt.Println(res)
}
