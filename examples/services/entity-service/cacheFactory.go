package main

import "time"

func createRedisCache(address string, expiration time.Duration) cache {
	client := createRedisClient(address)
	if client == nil {
		return &noneCache{}
	}
	return &redisClient{client: client, address: address, expiration: expiration}
}
