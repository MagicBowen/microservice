package main

import (
	"context"
	"log"
	"net"

	api "../api"
	"google.golang.org/grpc"
)

const (
	port = ":8899"
)

// server is used to implement interface.EntityServer.
type userEntityServer struct{}

func (s *userEntityServer) GetUser(ctx context.Context, req *api.UserRequest) (*api.UserRsp, error) {
	log.Printf("Get user: %v", req.Id)
	return &api.UserRsp{Name: "Bowen"}, nil
}

func (s *userEntityServer) AddUser(ctx context.Context, user *api.UserInfo) (*api.StatusMsg, error) {
	log.Printf("Add user: %v", user)
	return &api.StatusMsg{Code: 0}, nil
}

func (s *userEntityServer) UpdateUser(ctx context.Context, user *inf.UserInfoMsg) (*api.StatusMsg, error) {
	log.Printf("Update user: %v", user)
	return &api.StatusMsg{Code: 0}, nil
}

func (s *userEntityServer) DeleteUser(ctx context.Context, req *api.UserRequest) (*api.StatusMsg, error) {
	log.Printf("Delete user: %v", Req.Id)
	return &api.StatusMsg{Code: 0}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &userEntityServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
