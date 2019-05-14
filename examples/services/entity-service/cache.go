package main

import (
	"context"
	"errors"
)

type cache interface {
	set(ctx context.Context, key string, value string) error
	get(ctx context.Context, key string) (string, error)
	del(ctx context.Context, key string) error
}

type noneCache struct{}

func (nc *noneCache) set(ctx context.Context, key string, value string) error {
	return nil
}

func (nc *noneCache) get(ctx context.Context, key string) (string, error) {
	return "", errors.New("None cache for storing")
}

func (nc *noneCache) del(ctx context.Context, key string) error {
	return nil
}
