package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type entityCache struct {
	keyPrefix string
	cache     cache
}

func (uc *entityCache) getKey(id int) string {
	return fmt.Sprintf("%s/%d", uc.keyPrefix, id)
}

func (uc *entityCache) set(ctx context.Context, id int, entity interface{}) error {
	value, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	return uc.cache.set(ctx, uc.getKey(id), string(value))
}

func (uc *entityCache) get(ctx context.Context, id int, entity interface{}) error {
	value, err := uc.cache.get(ctx, uc.getKey(id))
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(value), entity)
}

func (uc *entityCache) del(ctx context.Context, id int) error {
	return uc.cache.del(ctx, uc.getKey(id))
}

func createEntityCache(prefix string, cache cache) *entityCache {
	return &entityCache{keyPrefix: prefix, cache: cache}
}
