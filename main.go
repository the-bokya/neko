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

// TODO: refactor
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
		if len(args) < 3 {
			fmt.Println("Please provide an argument for setup. Valid arguments:")
			fmt.Println("1. config")
			fmt.Println("2. images")
			return
		}
		switch args[2] {
		case "config":
			setup.Init(defaultConfigDir)
		case "images":
			// TODO: image setup
		default:
			fmt.Println("Please provide a valid argument for setup. Valid arguments:")
			fmt.Println("1. config")
			fmt.Println("2. images")
			return

		}
	case "serve":
		serve(app)
	default:
		if len(args) < 2 {
			fmt.Println("Please provide a valid argument. Valid arguments:")
			fmt.Println("1. setup")
			fmt.Println("2. serve")
			return
		}

	}

}
