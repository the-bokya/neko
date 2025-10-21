package setup

import (
	"embed"
	"os"
	"path/filepath"
)

const VMImagePath string = "/tmp/vm_image.qcow2"
const BaseImagePath string = "/tmp/base_image.qcow2"
const EtcPath string = "/etc/neko"

type Setup struct {
	Config           Config
	defaultConfigDir *embed.FS
}

func Init(defaultConfigDir embed.FS) error {
	setupData := &Setup{}
	setupData.defaultConfigDir = &defaultConfigDir
	if err := setupData.InitEtcDir(); err != nil {
		return err
	}
	return nil
}

// copy over the config to /etc/neko
func (setup *Setup) InitEtcDir() error {
	isPresent, err := isFilePresent(EtcPath)
	if isPresent {
		return nil
	}
	if err != nil {
		return err
	}
	if err = os.MkdirAll(EtcPath, 0755); err != nil {
		return err
	}
	configData, err := setup.defaultConfigDir.ReadFile("etc/config.json")
	if err != nil {
		return err
	}
	if err = os.WriteFile(filepath.Join(EtcPath, "config.json"), configData, 0644); err != nil {
		return err
	}
	return nil
}
