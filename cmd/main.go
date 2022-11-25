package main

import (
	"context"
	"grpc/cmd/config"
	"grpc/cmd/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"

	productPb "grpc/cmd/pb"

	"google.golang.org/grpc"
)

func main() {

	netListen, err := net.Listen("tcp", ":9090")
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
	go func ()  {
		if err := grpcServer.Serve(netListen); err != nil {
			log.Fatalf("Failed to serve %v", err.Error())
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		":9090",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to dial server : %s", err)
	}

	gwmux := runtime.NewServeMux()

	if err = productPb.RegisterProductServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalf("Failed to Register Gateway : %s", err)
	}

	gwServer := &http.Server{
		Addr:    ":9091",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on :9091")
	log.Fatalln(gwServer.ListenAndServe())
}
