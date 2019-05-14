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

func (db *mongoDB) findOne(ctx context.Context, filter interface{}, value interface{}) error {
	return db.collection.FindOne(ctx, filter).Decode(value)
}
func (db *mongoDB) insertOne(ctx context.Context, value interface{}) error {
	_, err := db.collection.InsertOne(ctx, value)
	return err
}
func (db *mongoDB) updateOne(ctx context.Context, filter interface{}, updateValue interface{}) error {
	_, err := db.collection.UpdateOne(ctx, filter, updateValue)
	return err
}

func (db *mongoDB) delete(ctx context.Context, filter interface{}) error {
	_, err := db.collection.DeleteMany(ctx, filter)
	return err
}
