/*
Package discovery for service discovery;
Usage:
	discovery := NewDiscovery().Follow("service1").Follow("service2").From(endpoints, path)
	defer discovery.Stop()
	serviceInstance := discovery.getInstanceOf("service1", discovery.RANDOM)
*/
package discovery

import (
	"context"
	"errors"
	"log"

	"go.etcd.io/etcd/clientv3"
)

type service struct {
	key       string
	instances map[string]*instance
	stop      chan interface{}
	client    *clientv3.Client
}

func newService(key string, client *clientv3.Client) *service {
	return &service{key: key,
		instances: make(map[string]*instance),
		stop:      make(chan interface{}),
		client:    client,
	}
}

func (s *service) addInstance(it *instance) {
	s.instances[it.getKey()] = it
}

func (s *service) removeInstance(it *instance) {
	delete(s.instances, it.getKey())
}

func (s *service) getInstance(algoType LBType) (string, error) {
	if len(s.instances) == 0 {
		return "", errors.New("None available instance of service")
	}
	for _, it := range s.instances {
		return it.getAddress(), nil
	}
	return "", errors.New("Internal error")
}

func (s *service) getAllInstances() ([]string, error) {
	if len(s.instances) == 0 {
		return nil, errors.New("None available instance of service")
	}
	var instances []string
	for _, it := range s.instances {
		instances = append(instances, it.getAddress())
	}
	return instances, nil
}

func (s *service) getResolver() (*resolver, error) {
	return newResolver(s), nil
}

func (s *service) dealEtcdEvents(events []*clientv3.Event) {
	for _, ev := range events {
		log.Printf("etcd: [%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		it := newInstance(string(ev.Kv.Value))
		switch ev.Type {
		case clientv3.EventTypePut:
			s.addInstance(it)
		case clientv3.EventTypeDelete:
			s.removeInstance(it)
		}
	}
}

func (s *service) following(ch clientv3.WatchChan) {
	for {
		select {
		case _ = <-s.stop:
			log.Printf("service(%s) has been shutdown", s.key)
			return
		case rsp, ok := <-ch:
			if !ok {
				log.Printf("following channel closed")
				return
			}
			s.dealEtcdEvents(rsp.Events)
		}
	}
}

func (s *service) fetchInstances() error {
	result, err := s.client.Get(context.Background(), s.key, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, ev := range result.Kvs {
		s.addInstance(newInstance(string(ev.Value)))
	}
	return nil
}

func (s *service) watchUpdate() clientv3.WatchChan {
	return s.client.Watch(context.Background(), s.key, clientv3.WithPrefix())
}

func (s *service) keepUpdate() {
	ch := s.watchUpdate()
	go s.following(ch)
}

func (s *service) follow() {
	s.fetchInstances()
	s.keepUpdate()
}

func (s *service) drop() {
	log.Printf("Stop service(%s) following", s.key)
	s.stop <- nil
}
