package discovery

import (
	"fmt"

	"google.golang.org/grpc/naming"
)

type resolver struct{ service *service }

func newResolver(service *service) *resolver {
	fmt.Printf("resolver created, service is %s\n", service.key)
	return &resolver{service: service}
}

func (re *resolver) Resolve(target string) (naming.Watcher, error) {
	return newWatcher(re.service), nil
}
