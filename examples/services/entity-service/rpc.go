package main

import (
	"context"
	"errors"
	"log"
	"net"

	api "github.com/magicbowen/microservice/examples/services/api"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/MagicBowen/microservice/examples/services/utils/tracing"
)

// server is used to implement api.EntityServer.
type userRPCServer struct {
	grpcServer *grpc.Server
	repo       *userRepo
	tracer     *tracing.ServiceTracer
}

func (s *userRPCServer) initServer() {
	s.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.tracer.OpenTracer())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(s.tracer.OpenTracer())),
	)
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

func (s *userRPCServer) startUp(address string, repo *userRepo, tracer *tracing.ServiceTracer) {
	s.repo = repo
	s.tracer = tracer
	s.initServer()
	s.run(address)
}

func (s *userRPCServer) GetUser(ctx context.Context, req *api.UserRequest) (*api.UserRsp, error) {
	log.Printf("Get user: %v", req.Id)
	if user := s.repo.getUserByID(int(req.Id)); user != nil {
		return &api.UserRsp{Name: user.Name}, nil
	}
	return nil, errors.New("Get none user!\n ")
}

func (s *userRPCServer) AddUser(ctx context.Context, user *api.UserInfoMsg) (*api.StatusMsg, error) {
	log.Printf("Add user: %v", user)
	if err := s.repo.createUser(createUserEntity(user)); err != nil {
		return &api.StatusMsg{Code: -1}, err
	}
	return &api.StatusMsg{Code: 0}, nil
}

func (s *userRPCServer) UpdateUser(ctx context.Context, user *api.UserInfoMsg) (*api.StatusMsg, error) {
	log.Printf("Update user: %v", user)
	if err := s.repo.updateUser(createUserEntity(user)); err != nil {
		return &api.StatusMsg{Code: -1}, err
	}
	return &api.StatusMsg{Code: 0}, nil
}

func (s *userRPCServer) DeleteUser(ctx context.Context, req *api.UserRequest) (*api.StatusMsg, error) {
	log.Printf("Delete user: %v", req.Id)
	if err := s.repo.deleteUser(int(req.Id)); err != nil {
		return &api.StatusMsg{Code: -1}, err
	}
	return &api.StatusMsg{Code: 0}, nil
}
