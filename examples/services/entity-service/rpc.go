package main

import (
	"context"
	"log"
	"net"

	api "github.com/magicbowen/microservice/examples/services/api"
	"google.golang.org/grpc"
)

// server is used to implement api.EntityServer.
type userRPCServer struct {
	grpcServer *grpc.Server
	repo       *userRepo
}

func (s *userRPCServer) initServer() {
	s.grpcServer = grpc.NewServer()
	api.RegisterEntityServer(s.grpcServer, s)
}

func (s *userRPCServer) run(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Entity server running on address %s", address)
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *userRPCServer) startUp(address string, repo *userRepo) {
	s.repo = repo
	s.initServer()
	s.run(address)
}

func (s *userRPCServer) GetUser(ctx context.Context, req *api.UserRequest) (*api.UserRsp, error) {
	log.Printf("Get user: %v", req.Id)
	return &api.UserRsp{Name: "Bowen"}, nil
}

func (s *userRPCServer) AddUser(ctx context.Context, user *api.UserInfoMsg) (*api.StatusMsg, error) {
	log.Printf("Add user: %v", user)
	return &api.StatusMsg{Code: 0}, nil
}

func (s *userRPCServer) UpdateUser(ctx context.Context, user *api.UserInfoMsg) (*api.StatusMsg, error) {
	log.Printf("Update user: %v", user)
	return &api.StatusMsg{Code: 0}, nil
}

func (s *userRPCServer) DeleteUser(ctx context.Context, req *api.UserRequest) (*api.StatusMsg, error) {
	log.Printf("Delete user: %v", req.Id)
	return &api.StatusMsg{Code: 0}, nil
}
