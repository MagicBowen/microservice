package main

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func testMongoConnection(client *mongo.Client, expiration time.Duration) error {
	ctx, cancel := getDefaultTimeoutCtx(expiration)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf("connect to mongodb failed: %v", err)
		return err
	}
	log.Printf("connect to mongodb successful! ")
	return nil
}

func createMongoConnection(address string, expiration time.Duration) *mongo.Client {
	ctx, cancel := getDefaultTimeoutCtx(expiration)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		log.Printf("create mongo client failed : %v", err)
		return nil
	}
	return client
}

func createMongoClient(address string, expiration time.Duration) *mongo.Client {
	if client := createMongoConnection(address, expiration); client != nil {
		if err := testMongoConnection(client, expiration); err == nil {
			return client
		}
	}
	return nil
}

func getCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func createMongoDB(address string, dbName string, collectionName string, expiration time.Duration) *mongoDB {
	client := createMongoClient(address, expiration)
	if client == nil {
		log.Fatalf("create mongodb failed")
	}
	collection := getCollection(client, dbName, collectionName)
	return &mongoDB{client: client, collection: collection, expiration: expiration}
}
