package main

import (
	"fmt"
	"log"
	"net"

	"github.com/imartingraham/grpc_demo/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := api.Server{}
	grpcServer := grpc.NewServer()

	api.RegisterUserServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
