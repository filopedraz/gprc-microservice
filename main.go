package main

import (
	"github.com/filopedraz/grpc-microservice/api"
	"github.com/filopedraz/grpc-microservice/gapi"
)

func main() {
	go api.RunHTTPServer()
	gapi.RunGRPCServer()
}
