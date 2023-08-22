package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc_sample/details" // Adjust the import path
)

type server struct {
	details.UnimplementedDetailsServiceServer
}

func (s *server) GetDetails(ctx context.Context, req *details.DetailsRequest) (*details.DetailsResponse, error) {
	log.Println("triggered .....")
	log.Println(ctx)
	// Simulate fetching details from a database or source
	name := req.Name
	age := req.Age

	response := &details.DetailsResponse{
		CommonRes: &details.CommonResponse{
			Code:    "SUCCESS",
			Message: "Your request success",
			Lang:    "en",
		},
		Name:  "My name is " + name,
		Age:   age,
		Email: req.Email,
	}

	return response, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	srv := grpc.NewServer()
	details.RegisterDetailsServiceServer(srv, &server{})

	fmt.Println("Server listening on port 50051")
	if err := srv.Serve(listen); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
