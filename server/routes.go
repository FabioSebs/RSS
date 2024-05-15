package server

import "github.com/labstack/echo/v4"

func SetRoutes(server *echo.Echo) {
	v1 := server.Group("/v1/icct") //append middlewares here
	{
		moe := v1.Group("/moe")
		{
			moe.GET("/xml", GetMoeXML)
		}
	}

}
