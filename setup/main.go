package setup

import (
	"embed"
	"os"
	"path/filepath"
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

	// for images from which vms will boot
	os.MkdirAll(filepath.Join(EtcPath, "disks"), 0755)
	return nil
}
