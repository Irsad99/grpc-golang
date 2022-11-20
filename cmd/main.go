package main

import (
	"grpc/cmd/config"
	"grpc/cmd/service"
	"log"
	"net"

	productPb "grpc/pb"

	"google.golang.org/grpc"
)

func main() {

	netListen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err.Error())
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to DB %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	productService := service.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
