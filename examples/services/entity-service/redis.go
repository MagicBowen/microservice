package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type redisProxy interface {
	set(key string, value string) error
	get(key string) (string, error)
}

type redisClient struct {
	address    string
	client     *redis.Client
	expiration time.Duration
}

type redisNone struct {
	address    string
	expiration time.Duration
}

func (rc *redisNone) set(key string, value string) error {
	return nil
}

func (rc *redisNone) get(key string) (string, error) {
	errStr := fmt.Sprintf("None connection redis proxy (%s)", rc.address)
	return "", errors.New(errStr)
}

func (rc *redisClient) set(key string, value string) error {
	return rc.client.Set(key, value, rc.expiration).Err()
}

func (rc *redisClient) get(key string) (string, error) {
	value, err := rc.client.Get(key).Result()
	if err == redis.Nil {
		errStr := fmt.Sprintf("key (%s) does not exist!", key)
		return "", errors.New(errStr)
	} else if err != nil {
		return "", err
	} else {
		return value, nil
	}
}

func (rc *redisClient) testConnection() error {
	pong, err := rc.client.Ping().Result()
	if err != nil {
		log.Printf("ping to redis(%s) failed (%s)", rc.address, err)
		return err
	}
	log.Printf("ping to redis(%s) successful (%s)", rc.address, pong)
	return nil
}

func newRedisProxy(address string, expiration time.Duration) redisProxy {
	var rc redisClient
	rc.address = address
	rc.expiration = expiration
	rc.client = redis.NewClient(&redis.Options{Addr: address, Password: "", DB: 0})
	if err := rc.testConnection(); err != nil {
		log.Printf("connect to redis(%s) failed: %v", address, err)
		return &redisNone{address: address, expiration: expiration}
	}
	log.Printf("connect to redis(%s) successful!", address)
	return &rc
}
