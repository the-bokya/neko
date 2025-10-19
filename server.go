package main

import (
	"embed"
	"fmt"
	"neko/application"
	"neko/libvirtapi"
	"neko/setup"

	"github.com/labstack/echo/v4"
)

//go:embed etc
var defaultConfigDir embed.FS

type Out struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	e := echo.New()
	lv := &libvirtapi.Libvirt{}
	err := lv.Init()
	conn := lv.Conn
	app := &application.Application{}
	app.LibvirtConn = conn
	setup.Init(defaultConfigDir)

	if err != nil {
		fmt.Println(err.Error())
	}
	e.GET("/hostname", app.GetHostname)
	e.Start("0.0.0.0:8000")
}
