package datauri

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// Encoder converts image files to Base64 string if successfully encoded.
// If error occures, it returns error.
func Encoder(imagePath string) (string, error) {
	if err := exsits(imagePath); err != nil {
		return "", err
	}

	image, err := os.Open(imagePath)
	defer image.Close()

	info, err := image.Stat()
	if err != nil {
		return "", err
	}

	size := info.Size()
	rawData := make([]byte, size)

	if readSize, err := image.Read(rawData); err != nil {
		return "", err
	} else if readSize == 0 {
		return "", fmt.Errorf("seemed failure reading the image file at %d bytes: %s", readSize, imagePath)
	}

	return base64.StdEncoding.EncodeToString(rawData), nil
}

func exsits(imagePath string) error {
	absPath, _ := filepath.Abs(imagePath)

	if _, err := os.Stat(absPath); os.IsNotExist(err) == true {
		return fmt.Errorf("no such file or directory: %s", absPath)
	}

	return nil
}

func writable(savePath string) error {
	absPath, _ := filepath.Abs(savePath)

	if _, err := os.Stat(absPath); os.IsNotExist(err) == false {
		return fmt.Errorf("the same name file already exitsts: %s", absPath)
	}

	saveDir := filepath.Dir(absPath)

	if d, err := os.Stat(saveDir); os.IsNotExist(err) == true || d.IsDir() == false {
		return fmt.Errorf("No such diretory exists: %s", saveDir)
	}

	return nil
}
