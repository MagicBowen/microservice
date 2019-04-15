package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func createUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}

	log.Printf("create user %v\n", u)
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
	// users[id].Name = u.Name
	return c.JSON(http.StatusOK, u)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("delete user %d\n", id)
	// delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

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

func initHTTPServer(address string) {
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
	e.Logger.Fatal(e.Start(address))
}
