package main

import (
	"fmt"
	"golang-microservices-gRPC/internal/gRPC/impl"
	"golang-microservices-gRPC/internal/gRPC/service"
	"log"
	"net"
	"google.golang.org/grpc"
)
const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	netListener := getNetListener()
	gRPCServer := grpc.NewServer()
	repositoryServiceImpl := impl.NewRepositoryServiceGrpcImpl()
	service.RegisterRepositoryServiceServer(gRPCServer,repositoryServiceImpl)
	log.Print(CONN_TYPE+"services start in port:"+CONN_PORT)

	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatal("fail to serve: %s",err)
	}
}

func getNetListener() net.Listener {
	lis, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("failed to listen: %v",err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
