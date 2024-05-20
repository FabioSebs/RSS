package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func GetXML(c echo.Context) error {
	var (
		name     string = c.Param("filename")
		filename string = fmt.Sprintf("xml/%s.xml", name)
	)

	// Read the XML file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading XML file")
	}

	// Set the content type to application/xml
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXML)
	return c.String(http.StatusOK, string(data))
}
