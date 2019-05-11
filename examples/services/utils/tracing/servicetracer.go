package tracing

import (
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ServiceTracer for service
type ServiceTracer struct {
	tracer opentracing.Tracer
	logger LogFactory
}

// NewServiceTracer to generate a global tracer with logger for service
func NewServiceTracer(serviceName string, metricsType MetricsType) *ServiceTracer {
	zlogger, _ := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	zapLogger := zlogger.With(zap.String("service", serviceName))
	logger := NewLogFactory(zapLogger)
	metricsFactory := NewMetrics(metricsType)
	tracer := NewTracer(serviceName, metricsFactory, logger)
	return &ServiceTracer{tracer: tracer, logger: logger}
}
