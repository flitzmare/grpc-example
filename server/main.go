package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/flitzmare/grpc-example/model"
	grpc "google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	var serverModel server
	var port = ":3000"

	model.RegisterExampleServer(srv, serverModel)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}

	srv.Serve(listener)
}

type server struct {
	model.UnimplementedExampleServer
}

func (s server) Hello(ctx context.Context, in *model.Request) (*model.Response, error) {
	fmt.Println("Hello Devfest, it's gRPC Server!")
	return &model.Response{
		Id:   in.Id,
		Name: in.Name,
	}, nil
}
