package tracing

import (
	"github.com/uber/jaeger-lib/metrics"
	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"
	jprom "github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewMetrics(backend string) metrics.Factory {
	switch backend {
	case "expvar":
		return jexpvar.NewFactory(10) // 10 buckets for histograms
	case "prometheus":
		return jprom.New().Namespace(metrics.NSOptions{Name: "hotrod", Tags: nil})
	default:
		return nil
	}
}
