package main

import (
	"context"
	"log"
	"time"

	api "github.com/magicbowen/microservice/examples/services/api"
	"google.golang.org/grpc"
)

type RPC struct {
	cc *grpc.ClientConn
	ec api.EntityClient
}

func (client *RPC) initial(address string) error {
	var err error
	client.cc, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect %s error: %v", address, err)
		return err
	}
	client.ec = api.NewEntityClient(client.cc)
	log.Printf("RPC client initialed successful\n")
	return nil
}

func (client *RPC) release() {
	if client.cc != nil {
		client.cc.Close()
		log.Printf("RPC client released successful\n")
	}
}

func (client *RPC) getUser(id int32) *user {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	u, err := client.ec.GetUser(ctx, &api.UserRequest{Id: id})
	if err != nil {
		log.Fatalf("get user error: %v", err)
	}
	return &user{ID: int(id), Name: u.Name}
}

var (
	rpc RPC
)
