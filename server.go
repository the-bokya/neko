package main

import (
	"neko/application"
	"neko/libvirtapi"
	"neko/setup"

	"github.com/labstack/echo/v4"
)

type Out struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func serve() error {
	lv := &libvirtapi.Libvirt{}
	err := lv.Init()
	if err != nil {
		return err
	}
	app := &application.Application{}
	app.Libvirt = lv
	appSetup := &setup.Setup{}
	appSetup.ReadConfig()
	app.VMConfig = &appSetup.Config

	e := echo.New()
	n := e.Group("/new")
	n.POST("/vm", app.CreateVM)
	e.Start("0.0.0.0:8000")
	return nil
}
