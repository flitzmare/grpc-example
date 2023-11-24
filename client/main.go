package main

import (
	"context"
	"fmt"
	"log"

	"github.com/flitzmare/grpc-example/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := model.NewExampleClient(conn)
	response, err := client.Hello(context.TODO(), &model.Request{
		Id:   1,
		Name: "Devfest",
	})
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("Hello from Client!\nID: %d\nName: %s\n", response.Id, response.Name)
}
