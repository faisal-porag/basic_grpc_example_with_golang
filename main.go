package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"

	"grpc_sample/details" // Adjust the import path
)

type server struct {
	details.UnimplementedDetailsServiceServer
}

func (s *server) GetDetails(ctx context.Context, req *details.DetailsRequest) (*details.DetailsResponse, error) {
	log.Println("triggered .....")
	// Simulate fetching details from a database or source
	name := req.Name
	age := req.Age

	response := &details.DetailsResponse{
		Code:    "SUCCESS",
		Message: "Your request is success",
		Lang:    "en",
		Data: &details.CommonResponse{
			Name:  "My name is " + name,
			Age:   age,
			Email: req.Email,
		},
	}

	return response, nil
}

func StartGRPCServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	srv := grpc.NewServer()
	details.RegisterDetailsServiceServer(srv, &server{})

	fmt.Println("Server listening on port 50051")
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
