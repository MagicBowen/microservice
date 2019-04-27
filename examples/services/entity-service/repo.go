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

type userRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

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

func (repo *userRepo) createUser(user *userEntity) error {
	ctx, cancel := getDefaultTimeoutCtx()
	defer cancel()
	_, err := repo.collection.InsertOne(ctx, bson.M{"id": user.ID, "name": user.Name})
	return err
}

func (repo *userRepo) getUserByID(id int) *userEntity {
	ctx, cancel := getDefaultTimeoutCtx()
	defer cancel()
	var user userEntity
	if err := repo.collection.FindOne(ctx, bson.M{"id": id}).Decode(&user); err != nil {
		log.Printf("find user by id(%d) failed: %v", id, err)
		return nil
	}
	return &user
}

func (repo *userRepo) updateUser(user *userEntity) error {
	ctx, cancel := getDefaultTimeoutCtx()
	defer cancel()
	_, err := repo.collection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": bson.M{"name": user.Name}})
	return err
}

func (repo *userRepo) deleteUser(id int) error {
	ctx, cancel := getDefaultTimeoutCtx()
	defer cancel()
	_, err := repo.collection.DeleteMany(ctx, bson.M{"id": id})
	return err
}

func createRepo(address string) *userRepo {
	repo := &userRepo{}
	repo.initial(address)
	return repo
}
