package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerOptions struct {
	Port string
}

// DefaultServerOptions returns the default configuration options for the server.
func DefaultServerOptions() *ServerOptions {
	return &ServerOptions{
		Port: "8080",
	}
}

// NewServer creates a new Echo server with the provided options.
func NewServer(options *ServerOptions) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
