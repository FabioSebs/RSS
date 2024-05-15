package server

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func GetMoeXML(c echo.Context) error {
	// Read the XML file
	data, err := ioutil.ReadFile("rss_feed.xml")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading XML file")
	}

	// Set the content type to application/xml
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXML)
	return c.String(http.StatusOK, string(data))
}
