package main

import (
	"context"
	"log"
	"net"
	productv1 "shop/api/product/v1"
	"shop/internal/server"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Transfer for transaction
func transfer(ctx context.Context, pool *pgxpool.Pool) error {

}

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

	reflection.Register(grpcServ)

	log.Printf("Server started on 50051")

	log.Fatal(grpcServ.Serve(lis))
}
