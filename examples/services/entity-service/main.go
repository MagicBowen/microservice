package main

import (
	"context"
	"log"
	"net"

	api "github.com/magicbowen/microservice/examples/services/api"
	"google.golang.org/grpc"
)

// server is used to implement interface.EntityServer.
type userEntityServer struct{}

func (s *userEntityServer) GetUser(ctx context.Context, req *api.UserRequest) (*api.UserRsp, error) {
	log.Printf("Get user: %v", req.Id)
	return &api.UserRsp{Name: "Bowen"}, nil
}

func (s *userEntityServer) AddUser(ctx context.Context, user *api.UserInfoMsg) (*api.StatusMsg, error) {
	log.Printf("Add user: %v", user)
	return &api.StatusMsg{Code: 0}, nil
}

func (s *userEntityServer) UpdateUser(ctx context.Context, user *api.UserInfoMsg) (*api.StatusMsg, error) {
	log.Printf("Update user: %v", user)
	return &api.StatusMsg{Code: 0}, nil
}

func (s *userEntityServer) DeleteUser(ctx context.Context, req *api.UserRequest) (*api.StatusMsg, error) {
	log.Printf("Delete user: %v", req.Id)
	return &api.StatusMsg{Code: 0}, nil
}

const (
	port = ":8899"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Entity server running on port %s", port)
	s := grpc.NewServer()
	api.RegisterEntityServer(s, &userEntityServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
