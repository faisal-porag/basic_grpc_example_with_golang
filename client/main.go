package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
	"grpc_sample/utils"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpc_sample/details"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:5005", grpc.WithInsecure())
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
	fmt.Printf("Name: %s\n", response.Data.Name)
	fmt.Printf("Age: %d\n", response.Data.Age)
	fmt.Printf("Email: %s\n", response.Data.Email)

	fmt.Printf("=========================================")
	fmt.Println()

	GetDetailsInfoWithAuth()
}

func GetDetailsInfoWithAuth() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:5005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Attach the JWT token as a Bearer token to the gRPC request
	token := generateServiceToken()
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)

	// Create a new client
	client := details.NewDetailsServiceClient(conn)

	// Prepare the request
	req := &details.DetailsRequest{
		Name:  "Faisal Porag",
		Age:   30,
		Email: "john@example.com",
	}

	// Call the GetDetails RPC
	response, err := client.GetDetailsWithAuthorization(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetDetails: %v", err)
	}

	fmt.Printf("%v\n", response)

	// Print the response
	fmt.Printf("Name: %s\n", response.Data.Name)
	fmt.Printf("Age: %d\n", response.Data.Age)
	fmt.Printf("Email: %s\n", response.Data.Email)
}

func generateServiceToken() string {
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &utils.ServiceAuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(utils.JwtSecretKey)
	if err != nil {
		return ""
	}

	return tokenString
}
