package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:55001")
	if err != nil {
		log.Fatalf("Error listening: %v\n", err)
		return
	}

	client := grpc.ClientConn()
}
