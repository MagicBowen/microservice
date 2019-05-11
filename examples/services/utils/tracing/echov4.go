package tracing

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
)

// Middleware returns a new Echo middleware handler for tracing
// requests and reporting errors.
func EchoMiddleware(tracer *ServiceTracer) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		m := &middleware{
			tracer:  tracer,
			handler: h,
		}
		return m.handle
	}
}

type middleware struct {
	handler echo.HandlerFunc
	tracer  *ServiceTracer
}

func (m *middleware) handle(c echo.Context) error {
	req := c.Request()
	opname := req.Method + ": " + c.Path()
	req, ht := nethttp.TraceRequest(m.tracer.OpenTracer(), req, nethttp.OperationName(opname))
	defer ht.Finish()
	return m.handler(c)
}
