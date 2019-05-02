package discovery

import (
	"log"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/naming"
)

type watcher struct {
	service       *service
	isInitialized bool
}

func newWatcher(service *service) *watcher {
	return &watcher{service: service}
}

func (w *watcher) Close() {
}

func (w *watcher) getInitializedUpates() ([]*naming.Update, bool) {
	err := w.service.fetchInstances()
	if err != nil {
		return nil, false
	}
	var addrs []string
	addrs, err = w.service.getAllInstances()
	if err != nil || len(addrs) == 0 {
		return nil, false
	}
	updates := make([]*naming.Update, len(addrs))
	for i := range addrs {
		updates[i] = &naming.Update{Op: naming.Add, Addr: addrs[i]}
	}
	return updates, true
}

func (w *watcher) getWatchUpdates() ([]*naming.Update, error) {
	ch := w.service.watchUpdate()

	for rsp := range ch {
		for _, ev := range rsp.Events {
			log.Printf("watch etcd: [%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			switch ev.Type {
			case clientv3.EventTypePut:
				return []*naming.Update{{Op: naming.Add, Addr: string(ev.Kv.Value)}}, nil
			case clientv3.EventTypeDelete:
				return []*naming.Update{{Op: naming.Delete, Addr: string(ev.Kv.Value)}}, nil
			}
		}
	}
	return nil, nil
}

func (w *watcher) Next() ([]*naming.Update, error) {
	if !w.isInitialized {
		w.isInitialized = true
		updates, ok := w.getInitializedUpates()
		if ok {
			return updates, nil
		}

	}
	return w.getWatchUpdates()
}
