package main

import (
	"context"
	"log"
	"time"

	"github.com/MagicBowen/microservice/examples/services/utils/discovery"
	api "github.com/magicbowen/microservice/examples/services/api"
	"google.golang.org/grpc"
)

var (
	etcdEndPoints = []string{"localhost:32773", "localhost:32771", "localhost:32769"}
)

const (
	servicePath       = "services"
	entityServiceName = "entity-service"
)

func main() {
	d, _ := discovery.NewDiscovery(etcdEndPoints, servicePath)
	defer d.Stop()

	r, err := d.Resolver(entityServiceName)
	if err != nil {
		log.Fatalf("Discovery initial resolver for gRPC failed: %v", err)
	}
	b := grpc.RoundRobin(r)

	client, err1 := grpc.Dial("", grpc.WithInsecure(), grpc.WithBalancer(b))
	if err1 != nil {
		log.Fatalf("connect with load balance error: %v", err1)
	}
	ec := api.NewEntityClient(client)

	time.Sleep(5 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	u, err2 := ec.GetUser(ctx, &api.UserRequest{Id: 1})
	if err2 != nil {
		log.Fatalf("get user error: %v", err2)
	}
	log.Printf("get user(%v) successful", u)
}
