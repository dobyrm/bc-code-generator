package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"math/rand"
	"time"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, World!")
	})

	e.GET("/rand", func(c echo.Context) error {
		response := struct {
			Status string `json:"status"`
			Data   struct {
				Rand int `json:"rand"`
			} `json:"data"`
		}{
			Status: "OK",
		}
		response.Data.Rand = Rand()

		return c.JSON(http.StatusOK, response)
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func Rand() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(101);
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
