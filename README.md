##### Basic gRPC(Remote Procedure Call) example with Golang

gRPC is a high-performance, open-source universal RPC (Remote Procedure Call) framework developers 
use to build highly scalable and distributed systems. It uses Protobuf (protocol buffers) as its 
interface definition language, which allows for a reliable way to define services and message types.
Microservices architecture is a way of designing software applications as suites of independently 
deployable services. It's a popular architecture for complex, evolving systems because it allows for 
scaling and allows teams to work on different services concurrently.

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
  
    ```
    sudo apt install -y protobuf-compiler
    ```
  
    ```shell
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        details/details.proto
    ```
    This will generate Go code for your service and messages in the specified directory.

    ****NOTE::**** If you're getting any error while execute `protoc` command then execute below commands on your terminal 

    ```shell
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
    ```
    ```shell
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    ```

  - `Implement gRPC Server:`
      - Create a Go file to implement the gRPC server.
      - Implement the methods defined in your service interface.
      - The methods should accept and return the protobuf message types you defined.
  
    ```shell
    go get google.golang.org/grpc
    ```
    
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

- `Client Implementation:`
    - Create a client to communicate with the gRPC server.
    - Import the generated protobuf package.
    - Use the generated client code to make calls to the gRPC methods.

    ```shell
    package main
    
    import (
        "context"
        "log"
    
        "google.golang.org/grpc"
        pb "path/to/your/generated/proto/package"
    )
    
    func main() {
        conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
        if err != nil {
            log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        client := pb.NewMyServiceClient(conn)
    
        // Create a context and call gRPC methods using the client
    }
    ```

- `Error Handling and Interceptors:`
    - Implement proper error handling in your server and client code.
    - Consider using gRPC interceptors for additional functionality like logging, authentication, etc.


> *****NOTE::***** Remember that gRPC and REST have different paradigms, so the way you structure your communication and handle errors might differ.
> Also, gRPC offers benefits like strong typing, automatic code generation, and support for streaming. 
> It's important to refactor your logic to fit the gRPC model effectively.


###### Project test instructions

- `For run server service:`
```shell
  make run_server
```

- `For run client service:`
```shell
make run_client
```


![gRPC_postman_response](https://github.com/faisal-porag/basic_grpc_example_with_golang/blob/dev_branch/files/grpc_postman_response.png?raw=true)

---

[Resource](https://protobuf.dev/getting-started/gotutorial/)

--- 

Thanks
