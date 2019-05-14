package main

import (
	"context"

	"github.com/MagicBowen/microservice/examples/services/utils/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"		
)


type spanFinish func ()

func traceProcess(
	ctx context.Context, 
	tracer *tracing.ServiceTracer, 
	action string, 
	key string, 
	value string,
	) (context.Context, spanFinish) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := tracer.OpenTracer().StartSpan(action + " " + key, opentracing.ChildOf(span.Context()))
		if action != "" {
			span.SetTag("param.action", action)
		}
		if key != "" {
			span.SetTag("param.key", key)
		}
		if value != "" {
			span.SetTag("param.value", value)
		}
		ext.SpanKindRPCClient.Set(span)
		ctx = opentracing.ContextWithSpan(ctx, span)
		return ctx, span.Finish
	}
	return ctx, func() {}
}