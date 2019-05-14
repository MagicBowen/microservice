package main

import (
	"context"
)

type entityDB interface {
	findOne(ctx context.Context, filter interface{}, value interface{}) error
	insertOne(ctx context.Context, value interface{}) error
	updateOne(ctx context.Context, filter interface{}, update interface{}) error
	delete(ctx context.Context, filter interface{}) error
}
