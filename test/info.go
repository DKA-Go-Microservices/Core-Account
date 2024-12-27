package main

import (
	"context"
	"fmt"
	"github.com/DKA-Go-Microservices/Core-Account/generated/proto/account/info"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Connect to the gRPC server on port 5051
	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())
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
	client := info.NewInfoClient(conn)

	/*res, err := client.Read(context.Background(), &info.ReadRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print the response message
	fmt.Println(res)*/

	res, err := client.Create(context.Background(), &info.InfoModel{
		FirstName: "Yovangga",
		LastName:  "Anandhika2",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print the response message
	fmt.Println(res)
}
