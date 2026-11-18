package server

import (
	"context"
	"log"
	productv1 "shop/api/product/v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	productv1.UnimplementedProductServiceServer

	products map[int64]*productv1.Product
}

func (s *Server) Init() {
	s.products = make(map[int64]*productv1.Product, 5)

	s.products[1] = &productv1.Product{Id: 1, Name: "Jopa"}
	s.products[2] = &productv1.Product{Id: 2, Name: "Pupa"}
	s.products[3] = &productv1.Product{Id: 3, Name: "lupa"}
	s.products[4] = &productv1.Product{Id: 4, Name: "Biba"}
	s.products[5] = &productv1.Product{Id: 5, Name: "Boba"}
}

func (s *Server) GetProduct(ctx context.Context, inReq *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	log.Println("Inside get product")
	val, ok := s.products[inReq.Id]
	if ok {
		return &productv1.GetProductResponse{Product: val}, nil
	} else {
		return nil, status.Error(codes.NotFound, "Product not found")
	}
}

func LoggingInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	log.Printf("Logging user\n")
	resp, err = handler(ctx, req)

	log.Printf("After handler\n")
	return resp, err
}

func ValidatorInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	log.Println("Do some validation")

	resp, err := handler(ctx, req)
	log.Println("After validation")
	return resp, err
}

func (s *Server) ListProducts(in *productv1.ListProductsRequest, stream grpc.ServerStreamingServer[productv1.Product]) error {
	for _, val := range s.products {
		if stream.Context().Err() != nil {
			return stream.Context().Err()
		}
		time.Sleep(1 * time.Second)

		if err := stream.Send(val); err != nil {
			return err
		}
	}

	return nil
}
