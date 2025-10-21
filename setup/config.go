package setup

import (
	"encoding/json"
	"io"
	"os"
)

const configPath string = "/etc/neko/config.json"
const localConfigPath string = "/etc/neko/config.json"

type Config struct {
	VMImages []VMImage `json:"vm_images"`
}

func (setup *Setup) ReadConfig() error {
	isPresent, err := isFilePresent(configPath)
	var configFile *os.File
	var configData []byte

	if isPresent {
		configFile, err = os.Open(configPath)
		defer configFile.Close()
		if err != nil {
			return err
		}
		configData, err = io.ReadAll(configFile)
	} else {
		return os.ErrNotExist
	}

	config, err := getConfig(configData)
	setup.Config = config
	return nil
}

func getConfig(configData []byte) (Config, error) {
	var config Config
	err := json.Unmarshal(configData, &config)
	return config, err
}
