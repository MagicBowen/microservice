package main

import (
	"time"

	"github.com/MagicBowen/microservice/examples/services/utils/registration"
	"github.com/MagicBowen/microservice/examples/services/utils/tracing"
)

const (
	mongoAddress      = "mongodb://mongodb:27017"
	dbName            = "microservice-example"
	collectionName    = "user"
	mongoOPExpiration = 2 * time.Second

	redisAddress    = "redis:6379"
	cacheExpiration = 5 * time.Second

	serviceAddress = ":8899"

	jaegerAgentAddress = "jaeger-agent:6831"
)

func main() {
	serviceTracer := tracing.NewServiceTracer("entity-service", jaegerAgentAddress, tracing.PROMETHEUS)
	serviceTracer.InfoLog("serviceTracer init OK")

	db := newTracingDB(serviceTracer, createMongoDB(mongoAddress, dbName, collectionName, mongoOPExpiration))
	cache := newTracingCache(serviceTracer, createRedisCache(redisAddress, cacheExpiration))
	repo := createUserRepo(db, cache)

	etcdEndpoints := []string{"etcd1:2379", "etcd2:2379", "etcd3:2379"}
	service := registration.NewService("entity-service").Address(serviceAddress).RegisterTo(etcdEndpoints, "services")
	defer service.Stop()

	(&userRPCServer{}).startUp(serviceAddress, repo, serviceTracer)
}
