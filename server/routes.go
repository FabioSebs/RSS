package server

import "github.com/labstack/echo/v4"

func SetRoutes(server *echo.Echo) {
	v1 := server.Group("/v1") //append middlewares here
	{
		rss := v1.Group("/rss")
		{
			rss.GET("/:filename", GetXML)
		}
	}

}
