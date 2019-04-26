package main

import (
	"context"
	"log"
	"net"
	"time"

	api "github.com/magicbowen/microservice/examples/services/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Printf("connect mongodb err: %v", err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("mongo service client err: %v!", err)
	}
	collection := client.Database("testing").Collection("numbers")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Printf("insert failed, %v", err)
	}
	id := res.InsertedID
	log.Printf("insert ok, id = %v", id)

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
