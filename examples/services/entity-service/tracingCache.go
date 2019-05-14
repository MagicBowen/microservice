package main

import (
	"context"
	"github.com/MagicBowen/microservice/examples/services/utils/tracing"
)

type tracingCache struct {
	cache cache
	tracer *tracing.ServiceTracer
}

func newTracingCache(tracer *tracing.ServiceTracer, cache cache) *tracingCache {
	return &tracingCache{cache : cache, tracer : tracer}
}

func (tc *tracingCache) traceRedis(
	ctx context.Context, 
	action string,
	key string, 
	value string,
	) (context.Context, spanFinish) {
	return traceProcess(ctx, tc.tracer, "redis " + action, key, value)
}

func (tc *tracingCache) set(ctx context.Context, key string, value string) error {
	_, spanFinish := tc.traceRedis(ctx, "set", key, value)
	defer spanFinish()
	return tc.cache.set(ctx, key, value)
}

func (tc *tracingCache) get(ctx context.Context, key string) (string, error) {
	_, spanFinish := tc.traceRedis(ctx, "get", key, "")
	defer spanFinish()
	return tc.cache.get(ctx, key)
}

func (tc *tracingCache) del(ctx context.Context, key string) error {
	_, spanFinish := tc.traceRedis(ctx, "del", key, "")
	defer spanFinish()
	return tc.cache.del(ctx, key)
}