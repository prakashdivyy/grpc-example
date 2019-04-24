package main

import (
  "fmt"
  "log"
	"net"
	
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-example/go/example"

)

type Server struct{}

// SayHello function implementation
func (s *Server) SayHello(context context.Context, in *example.HelloRequest) (*example.HelloResponse, error) {
	return &example.HelloResponse{Reply: "Hello " + in.Greeting + "!"}, nil
}

// DoAddition function implementation
func (s *Server) DoAddition(context context.Context, in *example.NumberRequest) (*example.NumberResponse, error) {
	return &example.NumberResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}

// DoSubtraction function implementation
func (s *Server) DoSubtraction(context context.Context, in *example.NumberRequest) (*example.NumberResponse, error) {
	return &example.NumberResponse{Result: in.FirstNumber - in.SecondNumber}, nil
}

func main() {
	// create a listener on TCP port 50050
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50050))
	if err != nil {
	  log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	example.RegisterExampleServiceServer(grpcServer, &Server{})

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
	  log.Fatalf("failed to serve: %s", err)
	}
}
