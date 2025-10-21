package setup

import (
	"embed"
)

const EtcPath string = "/etc/neko"

func Init(defaultConfigDir embed.FS) error {
	setupData := &Setup{}
	setupData.defaultConfigDir = &defaultConfigDir
	setupData.Init()
	return nil
}
func InitImages() error {
	setupData := &Setup{}
	setupData.InitImages()
	return nil
}
