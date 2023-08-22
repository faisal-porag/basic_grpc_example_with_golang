package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"grpc_sample/details"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := details.NewDetailsServiceClient(conn)

	// Prepare the request
	req := &details.DetailsRequest{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	// Call the GetDetails RPC
	response, err := client.GetDetails(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetDetails: %v", err)
	}

	fmt.Printf("%v\n", response)

	// Print the response
	fmt.Printf("Name: %s\n", response.Name)
	fmt.Printf("Age: %d\n", response.Age)
	fmt.Printf("Email: %s\n", response.Email)
}
