package mods

import (
	"io/ioutil"
	"os"

	"github.com/CCDirectLink/ccms/internal/utils"
)

// Extract h
func Extract(fileDesc *os.File) (string, error) {

	filePath, err := extract(fileDesc)

	if err != nil {
		return "", err
	}

	return filePath, nil
}

func extract(file *os.File) (string, error) {

	tempDir, err := ioutil.TempDir("", "ccms")

	if err != nil {
		return "", err
	}

	_, err = utils.Unzip(file.Name(), tempDir)

	if err != nil {
		return "", err
	}

	return tempDir, nil
}
