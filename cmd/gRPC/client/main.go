package main

import (
	"context"
	"fmt"
	"golang-microservices-gRPC/internal/gRPC/domain"
	"golang-microservices-gRPC/internal/gRPC/service"
	"google.golang.org/grpc"
	"log"
)

func main() {
	serverAddress := "localhost:3333"
	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatal("can't connect to server: %v", e.Error())
	}
	client := service.NewRepositoryServiceClient(conn)
	for i := range [10]int{} {
		repositoryModel := domain.Repository{
			Id:        int64(i),
			IsPrivate: true,
			Name:      string("Grpc-Demo"),
			UserId:    1245,
		}

		if responseMessage, e := client.Add(context.Background(), &repositoryModel); e != nil {
			panic(fmt.Sprintf("Was not able to insert Record %v", e))
		} else {
			fmt.Println("Record Inserted..")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}
