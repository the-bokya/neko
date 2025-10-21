// The libvirtxml package is an easy way to define the XML config for the VM
// It helps in avoiding interacting directly with raw XML
//

package libvirtapi

import (
	"github.com/google/uuid"
	"libvirt.org/go/libvirtxml"
)

type LibvirtConfig struct {
	Config *libvirtxml.Domain
}

func GenerateLibvirtConfig(name string) *LibvirtConfig {
	domcfg := &libvirtxml.Domain{Type: "kvm", Name: name, UUID: uuid.NewString()}
	config := &LibvirtConfig{}
	config.Config = domcfg
	config.Init()

	return config
}

// The idea is to programmatically define parameters like memory and vcpus through methods
// like VCPUs() and Memory()

func (config *LibvirtConfig) Memory(memory uint) {
	config.Config.Memory = &libvirtxml.DomainMemory{Unit: "KiB", Value: memory * 1024}
}
func (config *LibvirtConfig) VCPUs(num uint) {
	config.Config.VCPU = &libvirtxml.DomainVCPU{Value: num}
}
func (config *LibvirtConfig) AddFileDisk(file string) {
	disk := libvirtxml.DomainDisk{
		Device: "disk",
		Driver: &libvirtxml.DomainDiskDriver{
			Type: "qcow2",
			Name: "qemu"},
		Source: &libvirtxml.DomainDiskSource{
			File: &libvirtxml.DomainDiskSourceFile{
				File: file,
			},
		},
		Target: &libvirtxml.DomainDiskTarget{
			Dev: "vda",
			Bus: "virtio",
		},
		Address: &libvirtxml.DomainAddress{
			PCI: &libvirtxml.DomainAddressPCI{},
		},
	}
	config.Config.Devices.Disks = append(config.Config.Devices.Disks, disk)
}

func (config *LibvirtConfig) Network() {
	network := libvirtxml.DomainInterface{
		Source: &libvirtxml.DomainInterfaceSource{
			Network: &libvirtxml.DomainInterfaceSourceNetwork{
				Network: "default",
			},
		},
		Model: &libvirtxml.DomainInterfaceModel{Type: "virtio"},
	}
	config.Config.Devices.Interfaces = append(config.Config.Devices.Interfaces, network)
}

func (config *LibvirtConfig) Init() {
	config.Config.OS = &libvirtxml.DomainOS{Type: &libvirtxml.DomainOSType{Arch: "x86_64", Machine: "pc", Type: "hvm"}}
	if config.Config.Devices == nil {
		config.Config.Devices = &libvirtxml.DomainDeviceList{}
	}
	if config.Config.Devices.Disks == nil {
		config.Config.Devices.Disks = []libvirtxml.DomainDisk{}
	}
	if config.Config.Devices.Consoles == nil {
		config.Config.Devices.Consoles = []libvirtxml.DomainConsole{}
	}
	if config.Config.Devices.Interfaces == nil {
		config.Config.Devices.Interfaces = []libvirtxml.DomainInterface{}
	}
	config.Config.Devices.Consoles = append(config.Config.Devices.Consoles, libvirtxml.DomainConsole{Target: &libvirtxml.DomainConsoleTarget{Type: "serial"}})
}
