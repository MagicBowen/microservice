package main

import (
	"github.com/MagicBowen/microservice/examples/services/utils/discovery"
	"github.com/MagicBowen/microservice/examples/services/utils/tracing"
)

var (
	etcdEndPoints = []string{"etcd1:2379", "etcd2:2379", "etcd3:2379"}
)

const (
	servicePath        = "services"
	entityServiceName  = "entity-service"
	serviceAddress     = ":8866"
	jaegerAgentAddress = "jaeger-agent:6831"
)

func main() {
	serviceTracer := tracing.NewServiceTracer("http-service", jaegerAgentAddress, tracing.PROMETHEUS)
	serviceTracer.InfoLog("serviceTracer init OK")

	d, _ := discovery.NewDiscovery(etcdEndPoints, servicePath)
	defer d.Stop()

	err := rpc.initial(entityServiceName, d, serviceTracer)
	if err != nil {
		serviceTracer.FatalLog("gRPC init failed")
		return
	}
	defer rpc.release()

	initHTTPServer(serviceAddress, serviceTracer)
}
