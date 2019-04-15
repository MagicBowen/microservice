package main

import (
	"context"
	"errors"
	"fmt"
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

func (client *RPC) addUser(u *user) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.ec.AddUser(ctx, &api.UserInfoMsg{Id: int32(u.ID), Name: u.Name})
	if err != nil {
		log.Fatalf("add user error: %v", err)
		return err
	}
	if status.Code != 0 {
		errStr := fmt.Sprintf("add user error, status code: %d", status.Code)
		log.Fatalf(errStr)
		return errors.New(errStr)
	}
	return nil
}

func (client *RPC) updateUser(u *user) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.ec.UpdateUser(ctx, &api.UserInfoMsg{Id: int32(u.ID), Name: u.Name})
	if err != nil {
		log.Fatalf("update user error: %v", err)
		return err
	}
	if status.Code != 0 {
		errStr := fmt.Sprintf("update user error, status code: %d", status.Code)
		log.Fatalf(errStr)
		return errors.New(errStr)
	}
	return nil
}

func (client *RPC) deleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.ec.DeleteUser(ctx, &api.UserRequest{Id: int32(id)})
	if err != nil {
		log.Fatalf("delete user error: %v", err)
		return err
	}
	if status.Code != 0 {
		errStr := fmt.Sprintf("delete user error, status code: %d", status.Code)
		log.Fatalf(errStr)
		return errors.New(errStr)
	}
	return nil
}

var (
	rpc RPC
)

// type result interface{}

// func exec(method func(ctx context.Context) (result, error)) (result, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	return method(ctx)
// }

// func (client *RPC) getUser(id int32) *user {
// 	r, err := exec(func(ctx context.Context) (result, error) {
// 		u, err := client.ec.GetUser(ctx, &api.UserRequest{Id: id})
// 		if err != nil {
// 			log.Fatalf("get user error: %v", err)
// 			return nil, err
// 		}
// 		return &user{ID: int(id), Name: u.Name}, nil
// 	})
// 	if err != nil {
// 		return &user{ID: int(id), Name: "none"}
// 	}
// 	value, ok := r.(*user)
// 	if ok {
// 		return value
// 	}
// 	return &user{ID: int(id), Name: "none"}
// }
