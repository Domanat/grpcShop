package main

import (
	"log"
	"net"
	productv1 "shop/api/product/v1"
	"shop/internal/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	myServ := server.Server{}
	myServ.Init()

	grpcServ := grpc.NewServer(grpc.ChainUnaryInterceptor(
		server.LoggingInterceptor,
		server.ValidatorInterceptor,
	))

	productv1.RegisterProductServiceServer(grpcServ, &myServ)

	log.Printf("Server started on 50051")

	log.Fatal(grpcServ.Serve(lis))
}
