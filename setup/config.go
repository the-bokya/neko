package setup

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err)
		if err != nil {
			return err
		}
		configData, err = io.ReadAll(configFile)
	} else {
		if err != nil {
			return err
		} else {
			configData, err = setup.defaultConfigDir.ReadFile("etc/config.json")
			if err != nil {
				return err
			}
		}
	}

	config, err := getConfig(configData)
	fmt.Println(config)
	return nil
}

func getConfig(configData []byte) (Config, error) {
	var config Config
	err := json.Unmarshal(configData, &config)
	fmt.Println(config)
	return config, err
}
