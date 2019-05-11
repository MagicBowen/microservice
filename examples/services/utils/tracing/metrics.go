package tracing

import (
	"github.com/uber/jaeger-lib/metrics"
	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"
	jprom "github.com/uber/jaeger-lib/metrics/prometheus"
)

type MetricsType int

const (
	_ MetricsType = iota
	PROMETHEUS
	EXPVAR
)

func NewMetrics(backend MetricsType) metrics.Factory {
	switch backend {
	case EXPVAR:
		return jexpvar.NewFactory(10) // 10 buckets for histograms
	case PROMETHEUS:
		return jprom.New().Namespace(metrics.NSOptions{Name: "hotrod", Tags: nil})
	default:
		return nil
	}
}
