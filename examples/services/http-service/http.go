package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/MagicBowen/microservice/examples/services/utils/tracing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
   	httpCounters = prometheus.NewCounterVec(
   		prometheus.CounterOpts{
   			Name: "http_service_counters",
   			Help: "Number http requests",
   		},
   		[]string{"http_service"},
   	)	
)

func init() {
	prometheus.MustRegister(httpCounters)
}

func getContext(c echo.Context) context.Context {
	return c.Request().Context()
}

func createUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	err := rpc.addUser(getContext(c), u)
	if err != nil {
		log.Printf("create user %v failed\n", u)
		return c.NoContent(http.StatusForbidden)
	}
	log.Printf("create user %v success\n", u)
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("get user %d\n", id)
	log.Printf("req header = %v", c.Request().Header)
	user, err := rpc.getUser(getContext(c), int32(id))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	} 
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	u.ID, _ = strconv.Atoi(c.Param("id"))
	err := rpc.updateUser(getContext(c), u)
	if err != nil {
		log.Printf("update user %v failed\n", u)
		return c.NoContent(http.StatusForbidden)
	}
	log.Printf("update user %v success\n", u)
	return c.JSON(http.StatusOK, u)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := rpc.deleteUser(getContext(c), id)
	if err != nil {
		log.Printf("delete user %d failed\n", id)
		return c.NoContent(http.StatusForbidden)
	}
	log.Printf("delete user %d success\n", id)
	return c.String(http.StatusOK, "ok")
}

func getSiteName() string {
	host, err := os.Hostname()
	if err != nil {
		return "none"
	}
	return host
}

func health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func metrics(c echo.Context) error {
	log.Printf("handle prometheus collector request\n")
	promhttp.Handler().ServeHTTP(c.Response().Writer, c.Request())
	return nil
}

func metricMiddleware() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func (c echo.Context) error {
			defer func() {
				if c.Response().Status >= http.StatusBadRequest {
					httpCounters.With(prometheus.Labels{"http_service":"failedCount"}).Inc()
				} else {
					httpCounters.With(prometheus.Labels{"http_service":"sucessCount"}).Inc()
				}
				httpCounters.With(prometheus.Labels{"http_service":"totalCount"}).Inc()
			}()
			return h(c)
		}
	}
}

func home(c echo.Context) error {
	log.Printf("Home Page!\n")
	return c.String(http.StatusOK, "Welcome to "+getSiteName())
}

func initHTTPServer(address string, tracer *tracing.ServiceTracer) {
	e := echo.New()

	log.Printf("new server...\n")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(tracing.EchoMiddleware(tracer))
	e.Use(metricMiddleware())

	log.Printf("register middleware...\n")
	
	// Routes
	e.GET("/", home)
	e.GET("/health", health)
	e.GET("/metrics", metrics)

	log.Printf("set service handler...\n")

	// User Routes
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	log.Printf("server run on port 8866...\n")

	// Start server
	e.Logger.Fatal(e.Start(address))
}
