package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type redisClient struct {
	address    string
	expiration time.Duration
	client     *redis.Client
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

func (rc *redisClient) del(key string) error {
	return rc.client.Del(key).Err()
}
