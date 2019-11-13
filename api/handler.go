package api

import (
	"log"

	"github.com/imartingraham/grpc_demo/internal/model"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct{}

// FindUser finds user by id
func (s *Server) FindUser(ctx context.Context, in *UserMessage) (*UserMessage, error) {
	log.Printf("receive message %v", in.Id)
	user, err := model.FindUser(in.Id)
	if err != nil {
		return nil, err
	}
	return &UserMessage{Id: in.Id, FirstName: user.FirstName}, nil
}
