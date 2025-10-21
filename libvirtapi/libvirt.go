package libvirtapi

import (
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
	libvirtObject.Conn = l
	return nil

}

func (libvirtObject *Libvirt) CreateVM(name string, file string, numberOfVCPUs uint, memory uint) (*libvirt.Domain, error) {
	config := GenerateLibvirtConfig(name)
	config.AddFileDisk(file)
	config.Network()
	config.VCPUs(numberOfVCPUs)
	config.Memory(memory)
	xmlConfig, err := config.Config.Marshal()
	if err != nil {
		return nil, err
	}
	domain, err := libvirtObject.Conn.DomainDefineXML(xmlConfig)
	defer domain.Free()
	if err != nil {
		return nil, err
	}
	return domain, err
}
