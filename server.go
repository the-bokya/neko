package main

import (
	"github.com/labstack/echo/v4"
	"neko/application"
)

type Out struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func serve(app *application.Application) {
	e := echo.New()
	e.GET("/hostname", app.GetHostname)
	e.Start("0.0.0.0:8000")

}
