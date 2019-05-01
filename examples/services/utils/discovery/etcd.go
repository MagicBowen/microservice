package discovery

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func getEtcdClient(endpoints []string) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: 5 * time.Second})
}

func concatPathOf(path string) string {
	if path == "" {
		return ""
	}
	return fmt.Sprintf("/%s", path)
}
