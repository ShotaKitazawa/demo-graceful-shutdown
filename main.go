package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", func(c echo.Context) error {
		time.Sleep(20 * time.Second)
		return nil
	})
	e.GET("/health", func(c echo.Context) error {
		c.String(Healthcheck(), "")
		return nil
	})

	// Listen
	if err := http.ListenAndServe(":8080", e); err != nil {
		fmt.Fprintf(os.Stderr, "error serving webhook: %s", err)
		os.Exit(1)
	}
}

func Healthcheck() (code int) {
	_, err := os.Stat(os.Getenv("HOME") + "/.shutdown")
	if err != nil {
		return http.StatusOK
	}
	return http.StatusForbidden
}
