package main

import (
	"log"

	"github.com/go-redis/redis"
)

func testRedisConnection(client *redis.Client) error {
	if _, err := client.Ping().Result(); err != nil {
		return err
	}
	return nil
}

func createRedisClient(address string) *redis.Client {
	client := redis.NewClient(&redis.Options{Addr: address, Password: "", DB: 0})
	if err := testRedisConnection(client); err != nil {
		log.Printf("connect to redis(%s) failed (%v)", address, err)
		return nil
	}
	log.Printf("connect to redis(%s) successful!", address)
	return client
}
