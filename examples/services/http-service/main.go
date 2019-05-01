package main

import (
	"flag"
	"log"

	"github.com/MagicBowen/microservice/examples/services/utils/discovery"
)

var (
	logFile = flag.String("log", "output.log", "Log file name")
)

func main() {
	flag.Parse()
	// initLogger(logFile)

	d, _ := discovery.NewDiscovery([]string{"etcd1:2379", "etcd2:2379", "etcd3:2379"}, "services")
	defer d.Stop()
	d.Follow("entity-service")
	entityServiceAddress, err := d.InstanceOf("entity-service", discovery.Random)
	if err != nil {
		log.Fatalf("Get instance of entity service failed: %v", err)
	}

	err = rpc.initial(entityServiceAddress)
	if err != nil {
		log.Fatalf("gRPC init failed")
		return
	}
	defer rpc.release()

	initHTTPServer(":8866")
}
