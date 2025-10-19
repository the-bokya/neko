package setup

import (
	"errors"
	"os"
)

type VMImage struct {
	VMImagePath          string `json:"vm_image_path"`
	BaseImagePath        string `json:"base_image_path"`
	BaseImageDownloadURL string `json:"base_image_download_url"`
}

func (VMImage) setupVMImage() error {
	systemImageExists, err := isFilePresent(VMImagePath)

	if systemImageExists {
		return nil
	} else {
		if err != nil {
			return err
		}
		// createVMImage()
	}

	return err
}

// check if an attachable system image exists
func isFilePresent(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
