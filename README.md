# Go gPRC Microservice

Simple gRPC Server and Client in Go.

## Initialize the Project

```bash
go mod init github.com/filopedraz/grpc-microservice
go mod tidy # to install all the necessary dependencies
```

## Install the necessary dependencies for gRPC

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Generating Client and Server Code

```bash
make proto
```

## Run the Server

```bash
make server
```

## Roadmap

- [ ] Experiment with `evans` to interact with the server
- [ ] Explore [gPRC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) to expose gRPC services as RESTful APIs at the same time.
