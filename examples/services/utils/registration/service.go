/*
Package registration for service registration;
Usage:
	server, err := Registration.NewService("service1").Address("localhost:8080").RegisterTo(endpoints)
	defer server.stop()
*/
package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
)

const defaultPort = "8080"

type (
	// Service ...
	Service struct {
		name    string
		address string
		client  *clientv3.Client
		leaseID clientv3.LeaseID
		ttl     int64
		key     string
		stop    chan interface{}
	}
)

func NewService(name string) *Service {
	return &Service{name: name}
}

func (s *Service) Address(address string) *Service {
	s.address = address
	return s
}

func (s *Service) TTL(ttl int64) *Service {
	s.ttl = ttl
	return s
}

func (s *Service) RegisterTo(endpoints []string, pathPrefix string) *Service {
	s.saveParameters(endpoints, pathPrefix)
	go s.keepAlive()
	return s
}

func (s *Service) Stop() {
	s.stop <- nil
}

func (s *Service) revoke() {
	_, err := s.client.Revoke(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("etcd revoke failed: %v", err)
	}
	log.Printf("service:%s stopped\n", s.name)
}

func (s *Service) saveAddress() {
	addressParts := strings.Split(s.address, ":")
	var ip, port string
	if addressParts[0] == "" {
		ip = getLocalIP()
	} else {
		ip = addressParts[0]
	}
	if len(addressParts) == 1 || addressParts[1] == "" {
		port = defaultPort
	} else {
		port = addressParts[1]
	}
	s.address = fmt.Sprintf("%s:%s", ip, port)
}

func (s *Service) saveTTL() {
	if s.ttl == 0 {
		s.ttl = 5
	}
}

func (s *Service) saveClient(endpoints []string) {
	client, err := getEtcdClient(endpoints)
	if err != nil {
		log.Fatalf("create etcd client failed: %v", err)
	}
	s.client = client
}

func (s *Service) saveKey(pathPrefix string) {
	s.key = fmt.Sprintf("/%s/%s", pathPrefix, s.name)
}

func (s *Service) saveLease() {
	rsp, err := s.client.Grant(context.TODO(), s.ttl)
	if err != nil {
		log.Fatalf("etcd grant failed: %v", err)
	}
	_, err = s.client.Put(context.TODO(), s.key, s.address, clientv3.WithLease(rsp.ID))
	if err != nil {
		log.Fatalf("etcd put failed: %v", err)
	}
	s.leaseID = rsp.ID
}

func (s *Service) saveParameters(endpoints []string, pathPrefix string) {
	s.saveAddress()
	s.saveTTL()
	s.saveClient(endpoints)
	s.saveKey(pathPrefix)
	s.saveLease()
}

func (s *Service) keepAlive() {
	ch, err := s.client.KeepAlive(context.TODO(), s.leaseID)
	if err != nil {
		log.Fatalf("keep alive with etcd failed: %v", err)
	}

	for {
		select {
		case _ = <-s.stop:
			log.Printf("service has been shutdown")
			s.revoke()
			return
		case <-s.client.Ctx().Done():
			log.Printf("etcd server closed")
			return
		case ka, ok := <-ch:
			if !ok {
				log.Printf("keep alive channel closed")
				s.revoke()
				return
			}
			log.Printf("received reply from service: %s, ttl:%d", s.name, ka.TTL)
		}
	}
}

func getEtcdClient(endpoints []string) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: 5 * time.Second})
}

func main() {
	service := NewService("test").RegisterTo([]string{"localhost:9000"}, "services")
	defer service.Stop()
	time.Sleep(time.Second * 60)
}
