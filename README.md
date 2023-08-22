##### Basic gRPC(Remote Procedure Call) example with Golang

###### Some steps are given below:

- `Define Protobuf Messages and Services:`
    - Start by defining your service interface and messages in a `.proto` file using Protocol Buffers. 
        This includes specifying the request and response message types for each endpoint.
    - Define the service itself using the `service` keyword.

    ```shell
    syntax = "proto3";
    
    package myservice;
    
    message MyRequest {
        // Fields...
    }
    
    message MyResponse {
        // Fields...
    }
    
    service MyService {
        rpc MyMethod(MyRequest) returns (MyResponse);
    }
    ```

- `Generate gRPC Code:`
    - Use the `protoc` compiler to generate the gRPC code from your `.proto` file.
    - Run the following command to generate Go code:

    ```shell
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        details/details.proto
    ```
    This will generate Go code for your service and messages in the specified directory.


- `Implement gRPC Server:`
    - Create a Go file to implement the gRPC server.
    - Implement the methods defined in your service interface.
    - The methods should accept and return the protobuf message types you defined.
    
    ```shell
        package main
    
        import (
            "context"
            "log"
            "net"
        
            "google.golang.org/grpc"
            pb "path/to/your/generated/proto/package"
        )
        
        type server struct {
            pb.UnimplementedMyServiceServer
        }
        
        func (s *server) MyMethod(ctx context.Context, req *pb.MyRequest) (*pb.MyResponse, error) {
            // Implement your logic here
            // Create a MyResponse instance and return
        }
        
        func main() {
            lis, err := net.Listen("tcp", ":50051")
            if err != nil {
                log.Fatalf("failed to listen: %v", err)
            }
            s := grpc.NewServer()
            pb.RegisterMyServiceServer(s, &server{})
            if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
            }
        }
    ```
  
