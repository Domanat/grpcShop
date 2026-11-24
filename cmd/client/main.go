package main

import (
	"context"
	"errors"
	"log"
	"shop/internal/model"
	"shop/internal/storage/postgres"

	"github.com/jackc/pgx/v5"
)

func main() {
	// conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer conn.Close()

	// myClient := productv1.NewProductServiceClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	// defer cancel()

	// // For Unary
	// productResp, err := myClient.GetProduct(ctx, &productv1.GetProductRequest{Id: 2})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Result: %s\n", productResp.Product.Name)

	// For Stream
	// stream, err := myClient.ListProducts(ctx, &productv1.ListProductsRequest{})
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// for {
	// 	product, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	log.Printf("product : %s", product.Name)
	// }

	// Database work

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dsn := "postgres://app:secret@localhost:5432/products?sslmode=disable"

	pool, err := postgres.NewPool(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	prodRepository := postgres.NewProductRepository(pool)

	products := []model.Product{
		{Name: "Max", Price: 1111},
		{Name: "ASYA", Price: 2000},
		{Name: "LLL", Price: 3000},
		{Name: "KIRA", Price: 4000},
	}

	for _, p := range products {
		id, err := prodRepository.Create(ctx, p.Name, p.Price)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("Add %d successfully\n", id)
		}
	}

	id := 2

	product, err := prodRepository.GetById(ctx, int64(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("No actual data for this id %d", id)
		} else {
			log.Printf("Error %v", err)
		}
	} else {
		log.Printf("Get %s:%d succesfully", product.Name, product.Price)
	}

}
