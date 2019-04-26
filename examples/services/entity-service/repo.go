package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName         = "microservice-example"
	collectionName = "user"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	userRepo struct {
		client     *mongo.Client
		collection *mongo.Collection
	}
)

func getDefaultTimeoutCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func createClient(address string) (*mongo.Client, error) {
	ctx, cancel := getDefaultTimeoutCtx()
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI(address))
}

func (repo *userRepo) initClient(address string) {
	var err error
	if repo.client, err = createClient(address); err != nil {
		log.Fatalf("create mongodb err: %v", err)
	}
}

func (repo *userRepo) testConnection(address string) {
	ctx, cancel := getDefaultTimeoutCtx()
	defer cancel()
	if err := repo.client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("connect to mongodb failed: %v", err)
	} else {
		log.Printf("connect to mongodb(%s) successful! ", address)
	}
}

func (repo *userRepo) initCollection(client *mongo.Client) {
	repo.collection = client.Database(dbName).Collection(collectionName)
}

func (repo *userRepo) initial(address string) {
	repo.initClient(address)
	repo.testConnection(address)
	repo.initCollection(repo.client)
}

func (repo *userRepo) createUser(u *user) error {
	ctx, cancel := getDefaultTimeoutCtx()
	defer cancel()
	_, err := repo.collection.InsertOne(ctx, bson.M{"id": u.ID, "name": u.Name})
	return err
}

func createRepo(address string) *userRepo {
	repo := &userRepo{}
	repo.initial(address)
	return repo
}
