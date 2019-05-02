/*
Package discovery for service discovery;
Usage:
	d, _ := NewDiscovery([]string{"127.0.0.1:32773"}, "services")
	defer d.Stop()
	d.Follow("service1")
	d.Follow("service2")
	d.Drop("service2")
	serviceInstance, _ := d.InstanceOf("service1", discovery.Random)
*/
package discovery

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"go.etcd.io/etcd/clientv3"
)

type Discovery struct {
	pathPrefix string
	services   map[string]*service
	client     *clientv3.Client
}

type LBType int

const (
	_ LBType = iota
	Random
	RoundRobin
	RoundRobinWithWeight
	Hash
)

func NewDiscovery(endpoints []string, pathPrefix string) (*Discovery, error) {
	client, err := getEtcdClient(endpoints)
	if err != nil {
		log.Printf("create etcd client failed: %v", err)
		return nil, err
	}
	return &Discovery{pathPrefix: pathPrefix,
		services: make(map[string]*service),
		client:   client,
	}, nil
}

func (d *Discovery) getServiceKey(serviceName string) string {
	return fmt.Sprintf("%s%s", concatPathOf(d.pathPrefix), concatPathOf(serviceName))
}

func (d *Discovery) Follow(serviceName string) error {
	if strings.TrimSpace(serviceName) == "" {
		return errors.New("Illegal service name")
	}
	_, ok := d.services[serviceName]
	if ok {
		return nil
	}
	d.services[serviceName] = newService(d.getServiceKey(serviceName), d.client)
	d.services[serviceName].follow()
	return nil
}

func (d *Discovery) Drop(serviceName string) error {
	if strings.TrimSpace(serviceName) == "" {
		return errors.New("Illegal service name")
	}
	service, ok := d.services[serviceName]
	if !ok {
		return nil
	}
	service.drop()
	delete(d.services, serviceName)
	return nil
}

func (d *Discovery) Stop() error {
	log.Println("Stop discovery")
	for _, service := range d.services {
		service.drop()
	}
	return d.client.Close()
}

func (d *Discovery) InstanceOf(serviceName string, algoType LBType) (string, error) {
	if algoType != Random {
		return "", errors.New("Unsupported LB algorithm, only support Random now")
	}

	service, ok := d.services[serviceName]
	if !ok {
		return "", errors.New("Service " + serviceName + " has not been followed")
	}

	return service.getInstance(algoType)
}

func (d *Discovery) AllInstancesOf(serviceName string) ([]string, error) {
	service, ok := d.services[serviceName]
	if !ok {
		return nil, errors.New("Service " + serviceName + " has not been followed")
	}

	return service.getAllInstances()
}
