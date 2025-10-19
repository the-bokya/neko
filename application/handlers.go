package application

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string      `json:"status"` // Either ok or not okay
	Message interface{} `json:"message"`
}

func (app *Application) GetHostname(c echo.Context) error {
	hostname, err := app.LibvirtConn.GetHostname()
	if err != nil {
		return c.JSON(http.StatusOK, Response{Status: "Not Okay", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, Response{Status: "OK", Message: hostname})
}
