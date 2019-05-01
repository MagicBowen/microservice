package main

import (
	"time"

	"github.com/MagicBowen/microservice/examples/services/utils/registration"
)

const (
	mongoAddress      = "mongodb://mongodb:27017"
	dbName            = "microservice-example"
	collectionName    = "user"
	mongoOPExpiration = 2 * time.Second

	redisAddress    = "redis:6379"
	cacheExpiration = 5 * time.Second

	serviceAddress = ":8899"

	etcdEndpoints = []string{"etcd1:2379", "etcd2:2379", "etcd3:2379"}
)

func main() {
	db := createMongoDB(mongoAddress, dbName, collectionName, mongoOPExpiration)
	cache := createCache(redisAddress, cacheExpiration)
	repo := createUserRepo(db, cache)
	service := registration.NewService("http-service").Address(serviceAddress).RegisterTo(etcdEndpoints, "services")
	defer service.Stop()
	(&userRPCServer{}).startUp(serviceAddress, repo)
}
