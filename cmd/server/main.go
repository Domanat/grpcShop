package main

import (
	"context"
	"fmt"
	"log"
	db "shop/internal/repository/postgres/sqlc"
	"shop/internal/storage/postgres"
	"time"
)

func main() {
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// myServ := server.Server{}
	// myServ.Init()

	// grpcServ := grpc.NewServer(grpc.ChainUnaryInterceptor(
	// 	server.LoggingInterceptor,
	// 	server.ValidatorInterceptor,
	// ))
	// productv1.RegisterProductServiceServer(grpcServ, &myServ)

	// reflection.Register(grpcServ)

	// log.Printf("Server started on 50051")

	// log.Fatal(grpcServ.Serve(lis))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	dsn := "postgres://app:secret@localhost:5432/products?sslmode=disable"

	pool, err := postgres.NewPool(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	queries := db.New(pool)

	id, err := queries.CreateProduct(ctx, db.CreateProductParams{
		Name:  "Baobab",
		Price: 9999,
	})

	if err != nil {
		log.Fatal(err)
	}

	product, err := queries.GetProduct(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(product)

}
