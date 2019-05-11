package tracing

import (
	"net"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func example() {
	zlogger, _ := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	zapLogger := zlogger.With(zap.String("service", "customer"))
	logger := NewLogFactory(zapLogger)
	metricsFactory := NewMetrics("prometheus")
	tracer := NewTracer("customer", metricsFactory, logger)
	span := tracer.StartSpan("main process")
	span.Finish()
	servicePort := 8080
	serverAddr := net.JoinHostPort("0.0.0.0", strconv.Itoa(servicePort))
	logger.Bg().Info("Starting", zap.String("address", serverAddr))
}
