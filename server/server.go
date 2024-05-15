package server

import (
	"github.com/labstack/echo/v4"
)

// very basic server
func RunServer() {
	// Create a new Echo instance
	e := echo.New()

	SetRoutes(e)

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
