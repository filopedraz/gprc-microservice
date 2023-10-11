package gapi

import (
	"context"
	"log"
	"net"

	hello "github.com/filopedraz/grpc-microservice/pb"
	"google.golang.org/grpc"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func RunGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
