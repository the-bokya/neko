package main

import (
	"embed"
	"fmt"
	"neko/application"
	"neko/libvirtapi"
	"neko/setup"
	"os"
)

//go:embed etc
var defaultConfigDir embed.FS

func main() {
	lv := &libvirtapi.Libvirt{}
	err := lv.Init()
	if err != nil {
		fmt.Println(err.Error())
	}
	conn := lv.Conn
	app := &application.Application{}
	app.LibvirtConn = conn
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide an argument. Valid arguments:")
		fmt.Println("1. setup")
		fmt.Println("2. serve")
		return
	}
	switch args[1] {
	case "setup":
		setup.Init(defaultConfigDir)
	case "serve":
		serve(app)

	}

}
