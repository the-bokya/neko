package setup

import (
	"embed"
)

const VMImagePath string = "/tmp/vm_image.qcow2"
const BaseImagePath string = "/tmp/base_image.qcow2"
const EtcPath string = "/etc/neko"

func Init(defaultConfigDir embed.FS) error {
	setupData := &Setup{}
	setupData.defaultConfigDir = &defaultConfigDir
	setupData.Init()

	return nil
}
