package setup

import (
	"embed"
	"os"
	"path/filepath"
)

type Setup struct {
	Config           Config
	defaultConfigDir *embed.FS
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

func (setup *Setup) Init() error {
	if err := setup.InitEtcDir(); err != nil {
		return err
	}
	if err := setup.ReadConfig(); err != nil {
		return err
	}
	return nil
}
