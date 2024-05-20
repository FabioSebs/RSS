package server

import (
	"github.com/labstack/echo/v4"
)

// very basic server
func RunServer() {
	e := echo.New()
	SetRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
