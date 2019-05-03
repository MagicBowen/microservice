package discovery

import (
	"google.golang.org/grpc/naming"
)

type resolver struct{ service *service }

func newResolver(service *service) *resolver {
	return &resolver{service: service}
}

func (re *resolver) Resolve(target string) (naming.Watcher, error) {
	return newWatcher(re.service), nil
}
