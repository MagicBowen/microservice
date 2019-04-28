package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
	expiration time.Duration
}

func getDefaultTimeoutCtx(expiration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), expiration)
}

func (db *mongoDB) findOne(filter interface{}, value interface{}) error {
	ctx, cancel := getDefaultTimeoutCtx(db.expiration)
	defer cancel()
	return db.collection.FindOne(ctx, filter).Decode(value)
}
func (db *mongoDB) insertOne(value interface{}) error {
	ctx, cancel := getDefaultTimeoutCtx(db.expiration)
	defer cancel()
	_, err := db.collection.InsertOne(ctx, value)
	return err
}
func (db *mongoDB) updateOne(filter interface{}, updateValue interface{}) error {
	ctx, cancel := getDefaultTimeoutCtx(db.expiration)
	defer cancel()
	_, err := db.collection.UpdateOne(ctx, filter, updateValue)
	return err
}

func (db *mongoDB) delete(filter interface{}) error {
	ctx, cancel := getDefaultTimeoutCtx(db.expiration)
	defer cancel()
	_, err := db.collection.DeleteMany(ctx, filter)
	return err
}
