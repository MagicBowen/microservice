package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	api "github.com/magicbowen/microservice/examples/services/api"
	"google.golang.org/grpc"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//----------
// Logger
//----------

var (
	logFile = flag.String("log", "output.log", "Log file name")
)

func initLogger() {
	outfile, err := os.OpenFile(*logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}
	log.SetOutput(outfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

//----------
// type
//----------

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

//----------
// gRPC Client
//----------

const (
	entityServerAddress = "localhost:8899"
)

type clientRPC struct {
	cc *grpc.ClientConn
	ec api.EntityClient
}

func (client *clientRPC) initial() error {
	var err error
	client.cc, err = grpc.Dial(entityServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect %s error: %v", entityServerAddress, err)
		return err
	}
	client.ec = api.NewEntityClient(client.cc)
	log.Printf("RPC client initialed successful\n")
	return nil
}

func (client *clientRPC) release() {
	if client.cc != nil {
		client.cc.Close()
		log.Printf("RPC client released successful\n")
	}
}

func (client *clientRPC) getUser(id int32) *user {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// log.Printf("client ec is %v\n", client.ec)
	u, err := client.ec.GetUser(ctx, &api.UserRequest{Id: id})
	if err != nil {
		log.Fatalf("get user error: %v", err)
	}
	return &user{ID: int(id), Name: u.Name}
}

var (
	rpc clientRPC
)

//----------
// Http Api
//----------

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Http Handlers
//----------

func createUser(c echo.Context) error {
	log.Printf("create user %d\n", seq)
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("get user %d\n", id)
	user := rpc.getUser(int32(id))
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("update user %d\n", id)
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("delete user %d\n", id)
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

//----------
// Health Entrypoint
//----------

func getSiteName() string {
	host, err := os.Hostname()
	if err != nil {
		return "none"
	}
	return host
}

func home(c echo.Context) error {
	log.Printf("Home Page!\n")
	return c.String(http.StatusOK, "Welcome to "+getSiteName())
}

//----------
// Initial and setup
//----------

func main() {
	flag.Parse()
	// initLogger()

	err := rpc.initial()
	if err != nil {
		log.Fatalf("gRPC init failed")
		return
	}
	defer rpc.release()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", home)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	log.Printf("server run on port 8866...\n")

	// Start server
	e.Logger.Fatal(e.Start(":8866"))
}
