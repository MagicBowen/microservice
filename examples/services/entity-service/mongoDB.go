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

func (db *mongoDB) findOne(ctx context.Context, filter interface{}, value interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, db.expiration)
	defer cancel()
	return db.collection.FindOne(ctx, filter).Decode(value)
}
func (db *mongoDB) insertOne(ctx context.Context, value interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, db.expiration)
	defer cancel()	
	_, err := db.collection.InsertOne(ctx, value)
	return err
}
func (db *mongoDB) updateOne(ctx context.Context, filter interface{}, updateValue interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, db.expiration)
	defer cancel()
	_, err := db.collection.UpdateOne(ctx, filter, updateValue)
	return err
}

func (db *mongoDB) delete(ctx context.Context, filter interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, db.expiration)
	defer cancel()
	_, err := db.collection.DeleteMany(ctx, filter)
	return err
}
