package application

import (
	"neko/libvirtapi"
	"neko/setup"
)

type Application struct {
	VMConfig *setup.Config
	Libvirt  *libvirtapi.Libvirt
}
