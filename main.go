package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"grpc_sample/service_middleware"
	"grpc_sample/utils"
	"log"
	"net"
	"net/http"
	"time"

	"grpc_sample/details" // Adjust the import path
)

type server struct {
	details.UnimplementedDetailsServiceServer
}

func (s *server) GetDetails(ctx context.Context, req *details.DetailsRequest) (*details.DetailsResponse, error) {
	log.Println("triggered ...GetDetails...")
	// Simulate fetching details from a database or source
	name := req.Name
	age := req.Age

	response := &details.DetailsResponse{
		Code:    "SUCCESS",
		Message: "Your request is success",
		Lang:    "en",
		Data: &details.DataResponse{
			Name:  "My name is " + name,
			Age:   age,
			Email: req.Email,
		},
	}

	return response, nil
}

func (s *server) GetDetailsWithAuthorization(ctx context.Context, req *details.DetailsRequest) (*details.DetailsResponse, error) {
	log.Println("triggered ...GetDetailsWithAuthorization...")
	// Check if the request has valid authorization
	if err := service_middleware.CheckAuthorizationMiddleware(ctx); err != nil {
		return nil, err
	}

	// Simulate fetching details from a database or source
	name := req.Name
	age := req.Age

	response := &details.DetailsResponse{
		Code:    "SUCCESS",
		Message: "Your request is success",
		Lang:    "en",
		Data: &details.DataResponse{
			Name:  "My name is " + name,
			Age:   age,
			Email: req.Email,
		},
	}

	return response, nil
}

func StartGRPCServer() {
	listen, err := net.Listen("tcp", ":5005")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	srv := grpc.NewServer()
	details.RegisterDetailsServiceServer(srv, &server{})

	fmt.Println("Server listening on port 5005")
	go func() {
		if err := srv.Serve(listen); err != nil {
			fmt.Printf("Error: %v", err)
		}
	}()
}

func main() {
	// Start gRPC server
	StartGRPCServer()

	// *********************** ---HTTP SERVER--- ***********************

	// Start HTTP server
	http.HandleFunc("/api/rest/endpoint", handleRestRequest)
	http.HandleFunc("/get-token", handleCreateTokenRequest)
	fmt.Println("HTTP Server listening on port 8088")
	if err := http.ListenAndServe(":8088", nil); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func handleRestRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate fetching data for the JSON response
	data := map[string]interface{}{
		"message": "Hello from REST API!",
	}

	// Convert the data map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the JSON response
	_, _ = w.Write(jsonData)
}

func handleCreateTokenRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Simulate fetching data for the JSON response
	data := map[string]interface{}{
		"token": tokenString,
	}

	// Convert the data map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the JSON response
	_, _ = w.Write(jsonData)
}
