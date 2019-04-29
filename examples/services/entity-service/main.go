package main

import "time"

const (
	mongoAddress      = "mongodb://mongodb:27017"
	dbName            = "microservice-example"
	collectionName    = "user"
	mongoOPExpiration = 2 * time.Second

	redisAddress    = "redis:6379"
	cacheExpiration = 5 * time.Second

	serviceAddress = ":8899"
)

func main() {
	db := createMongoDB(mongoAddress, dbName, collectionName, mongoOPExpiration)
	cache := createCache(redisAddress, cacheExpiration)
	repo := createUserRepo(db, cache)
	(&userRPCServer{}).startUp(serviceAddress, repo)
}
