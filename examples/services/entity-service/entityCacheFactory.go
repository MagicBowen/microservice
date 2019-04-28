package main

import "time"

func createEntityCache(address string, expiration time.Duration) entityCache {
	client := createRedisClient(address)
	if client == nil {
		return &noneCache{}
	}
	return &redisCache{client: client, address: address, expiration: expiration}
}
