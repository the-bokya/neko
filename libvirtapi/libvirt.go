package libvirtapi

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

type Libvirt struct {
	Conn *libvirt.Connect
}

func (libvirtObject *Libvirt) Init() error {
	l, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return err
	}
	a, err := l.GetHostname()
	fmt.Println(a, "meow")
	libvirtObject.Conn = l
	return nil

}
