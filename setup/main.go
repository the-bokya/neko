package setup

import (
	"embed"
	"fmt"
)

const VMImagePath string = "/tmp/vm_image.qcow2"
const BaseImagePath string = "/tmp/base_image.qcow2"

type Setup struct {
	Config           Config
	defaultConfigDir *embed.FS
}

func Init(defaultConfigDir embed.FS) error {
	setupData := Setup{}
	setupData.defaultConfigDir = &defaultConfigDir
	err := setupData.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
