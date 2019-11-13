package main

import (
	"github.com/imartingraham/grpc_demo/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewUserClient(conn)
	response, err := c.FindUser(context.Background(), &api.UserMessage{Id: 1})
	if err != nil {
		log.Fatalf("error when calling FindUser: %s", err)
	}
	log.Printf("response from server: %s", response.FirstName)

}
