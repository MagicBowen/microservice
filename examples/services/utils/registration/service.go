/*
Package registration for service registration;
Usage:
	service := NewService("test").Address("127.0.0.1:8080").RegisterTo([]string{"localhost:2379"}, "services")
	defer service.Stop()
*/
package registration

import (
	"context"
	"fmt"
	"log"
	"strings"

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
	s.initParameters(endpoints, pathPrefix)
	go s.keepAlive()
	return s
}

func (s *Service) Stop() {
	log.Println("active stop service keepalive")
	s.stop <- nil
}

func (s *Service) revoke() {
	_, err := s.client.Revoke(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("etcd revoke failed: %v", err)
	}
	log.Printf("service:%s stopped\n", s.name)
}

func (s *Service) initAddress() {
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

func (s *Service) initTTL() {
	if s.ttl == 0 {
		s.ttl = 5
	}
}

func (s *Service) initClient(endpoints []string) {
	client, err := getEtcdClient(endpoints)
	if err != nil {
		log.Fatalf("create etcd client failed: %v", err)
	}
	s.client = client
}

func (s *Service) initKey(pathPrefix string) {
	s.key = fmt.Sprintf("%s%s%s", concatPathOf(pathPrefix), concatPathOf(s.name), concatPathOf(s.address))
}

func (s *Service) initLease() {
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

func (s *Service) initStopChan() {
	s.stop = make(chan interface{})
}

func (s *Service) initParameters(endpoints []string, pathPrefix string) {
	s.initAddress()
	s.initTTL()
	s.initClient(endpoints)
	s.initKey(pathPrefix)
	s.initLease()
	s.initStopChan()
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
			s.client.Close()
			return
		case <-s.client.Ctx().Done():
			log.Printf("etcd server closed")
			return
		case _, ok := <-ch:
			if !ok {
				log.Printf("keep alive channel closed")
				s.revoke()
				return
			}
		}
	}
}
