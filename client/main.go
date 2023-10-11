package main

import (
	"context"
	"fmt"

	proto "github.com/filopedraz/grpc-microservice/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		":50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	response, error := client.SayHello(context.Background(), &proto.HelloRequest{Name: "Filipe"})
	if error != nil {
		panic(error)
	}

	fmt.Println(response.Message)
}
