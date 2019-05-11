package tracing

import (
	"context"
	"fmt"

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

// InfoLog wrapper Logger.Info
func (st *ServiceTracer) InfoLog(msg string, fields ...interface{}) {
	st.logger.Bg().Info(fmt.Sprintf(msg, fields...))
}

// ErrorLog wrapper Logger.Info
func (st *ServiceTracer) ErrorLog(msg string, fields ...interface{}) {
	st.logger.Bg().Error(fmt.Sprintf(msg, fields...))
}

// FatalLog wrapper Logger.Info
func (st *ServiceTracer) FatalLog(msg string, fields ...interface{}) {
	st.logger.Bg().Fatal(fmt.Sprintf(msg, fields...))
}

// ContextLogger to get the Logger of context
func (st *ServiceTracer) ContextLogger(ctx context.Context) Logger {
	return st.logger.For(ctx)
}

// OpenTracer to get the tracer of opentracing
func (st *ServiceTracer) OpenTracer() opentracing.Tracer {
	return st.tracer
}
