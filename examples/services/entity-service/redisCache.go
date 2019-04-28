package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	address    string
	expiration time.Duration
	client     *redis.Client
}

func (rc *redisCache) set(key string, value string) error {
	return rc.client.Set(key, value, rc.expiration).Err()
}

func (rc *redisCache) get(key string) (string, error) {
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
