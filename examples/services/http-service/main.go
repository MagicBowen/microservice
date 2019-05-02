package main

import (
	"flag"
	"log"

	"github.com/MagicBowen/microservice/examples/services/utils/discovery"
)

var (
	logFile       = flag.String("log", "output.log", "Log file name")
	etcdEndPoints = []string{"etcd1:2379", "etcd2:2379", "etcd3:2379"}
)

const (
	servicePath       = "services"
	entityServiceName = "entity-service"
	serviceAddress    = ":8866"
)

func main() {
	flag.Parse()
	// initLogger(logFile)

	d, _ := discovery.NewDiscovery(etcdEndPoints, servicePath)
	defer d.Stop()

	err := rpc.initial(d, entityServiceName)
	if err != nil {
		log.Fatalf("gRPC init failed")
		return
	}
	defer rpc.release()

	initHTTPServer(serviceAddress)
}
