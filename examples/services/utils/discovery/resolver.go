package discovery

import (
	"google.golang.org/grpc/naming"
)

type resolver struct{ service *service }

func newResolver(service *service) *resolver {
	fmt.printf("resolver created, service is %s", service.key)
	return &resolver{service: service}
}

func (re *resolver) Resolve(target string) (naming.Watcher, error) {
	return newWatcher(re.service), nil
}
