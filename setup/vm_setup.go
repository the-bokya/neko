package setup

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"

	"github.com/schollz/progressbar/v3"
)

type VMStatus int

const (
	VMImageNotPresent VMStatus = iota
	VMImageDownloaded
)

var ErrImageDownload = errors.New("Error somewhere during image download")

type VMImage struct {
	Name                 string `json:"name"`
	BaseImageDownloadURL string `json:"base_image_download_url"`
	SHA256Sum            string `json:"sha256sum"`
	Status               VMStatus
}

func (image *VMImage) SetupVMImage() error {
	if image.Status == VMImageDownloaded {
		return nil
	}
	systemImageExists, err := image.doesSystemImageExist()

	if systemImageExists {
		return nil
	} else {
		if err != nil {
			return err
		}
		if err = os.MkdirAll(image.getBaseFolderPath(), 0644); err != nil {
			return err
		}
		for i := range 3 {
			if err = image.downloadVMImage(); err != nil {
				if errors.Is(err, ErrImageDownload) {
					if i < 2 {
						fmt.Println("There was an error downloading the image. Retrying.")
						continue
					}
					return err
				}
			} else {
				break
			}

		}
		return nil
	}
}

func (image *VMImage) downloadVMImage() error {
	baseImageDownloadStream, err := http.Get(image.BaseImageDownloadURL)
	defer baseImageDownloadStream.Body.Close()
	if err != nil {
		return err
	}
	baseImageFile, err := os.Create(image.getBaseImagePath())
	if err != nil {
		return err
	}
	bar := progressbar.DefaultBytes(-1, fmt.Sprintf("Downloading image for %s", image.Name))
	defer bar.Close()
	_, err = io.Copy(io.MultiWriter(baseImageFile, bar), baseImageDownloadStream.Body)
	if err != nil {
		return err
	}
	targetSHA256, err := hex.DecodeString(image.SHA256Sum)
	if err != nil {
		return err
	}
	baseImageFile.Close()
	hashesMatch, hashSum, err := image.verifySHA256Sum(image.getBaseImagePath(), targetSHA256)
	if err != nil {
		return err
	} else {
		fmt.Printf("Downloaded file hash: %x\n", hashSum)
		fmt.Printf("Actual hash provided in the config: %x\n", targetSHA256)
		if hashesMatch {
			fmt.Println("Successfully matched!")
		} else {
			fmt.Println("Not matching")
			os.Remove(image.getBaseImagePath())
			return ErrImageDownload
		}
	}
	return nil
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

func (image *VMImage) getBaseFolderPath() string {
	return filepath.Join(EtcPath, image.Name)
}
func (image *VMImage) getBaseImagePath() string {
	return filepath.Join(image.getBaseFolderPath(), "base_image")
}
func (image *VMImage) doesSystemImageExist() (bool, error) {
	exists, err := isFilePresent(image.getBaseImagePath())
	if exists {
		image.Status = VMImageDownloaded
		return true, nil
	}
	return exists, err
}

func (image *VMImage) verifySHA256Sum(filePath string, targetSHA256 []byte) (bool, []byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, []byte(""), err
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return false, []byte(""), err
	}
	hashSum := hash.Sum(nil)
	if slices.Equal(hashSum, targetSHA256) {
		return true, hashSum, nil
	} else {
		return false, hashSum, nil
	}
}
