package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/MagicBowen/microservice/examples/services/utils/discovery"
	api "github.com/magicbowen/microservice/examples/services/api"
	"google.golang.org/grpc"
)

type RPC struct {
	cc *grpc.ClientConn
	ec api.EntityClient
}

func (client *RPC) initial(d *discovery.Discovery, targetServiceName string) error {
	r, err := d.Resolver(targetServiceName)
	if err != nil {
		log.Fatalf("Discovery initial resolver for gRPC failed: %v", err)
	}
	b := grpc.RoundRobin(r)

	client.cc, err = grpc.Dial("", grpc.WithInsecure(), grpc.WithBalancer(b))
	if err != nil {
		log.Fatalf("connect with load balance error: %v", err)
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

func (client *RPC) getUser(id int32) (*user, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	u, err := client.ec.GetUser(ctx, &api.UserRequest{Id: id})
	if err != nil {
		errStr := fmt.Sprintf("get user error: %v", err)
		log.Printf(errStr)
		return nil, errors.New(errStr)
	}
	return &user{ID: int(id), Name: u.Name}, nil
}

func (client *RPC) addUser(u *user) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.ec.AddUser(ctx, &api.UserInfoMsg{Id: int32(u.ID), Name: u.Name})
	if err != nil {
		log.Printf("add user error: %v", err)
		return err
	}
	if status.Code != 0 {
		errStr := fmt.Sprintf("add user error, status code: %d", status.Code)
		log.Printf(errStr)
		return errors.New(errStr)
	}
	return nil
}

func (client *RPC) updateUser(u *user) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.ec.UpdateUser(ctx, &api.UserInfoMsg{Id: int32(u.ID), Name: u.Name})
	if err != nil {
		log.Printf("update user error: %v", err)
		return err
	}
	if status.Code != 0 {
		errStr := fmt.Sprintf("update user error, status code: %d", status.Code)
		log.Printf(errStr)
		return errors.New(errStr)
	}
	return nil
}

func (client *RPC) deleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.ec.DeleteUser(ctx, &api.UserRequest{Id: int32(id)})
	if err != nil {
		log.Printf("delete user error: %v", err)
		return err
	}
	if status.Code != 0 {
		errStr := fmt.Sprintf("delete user error, status code: %d", status.Code)
		log.Printf(errStr)
		return errors.New(errStr)
	}
	return nil
}

var (
	rpc RPC
)
