package main

import (
	"context"
	"log"
	productv1 "shop/api/product/v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	myClient := productv1.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	// For Unary
	productResp, err := myClient.GetProduct(ctx, &productv1.GetProductRequest{Id: 2})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Result: %s\n", productResp.Product.Name)

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

}
