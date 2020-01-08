package mods

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/CCDirectLink/ccms/internal/database"
)

// Download a mod by name
func Download(name string, dbType string) (*os.File, error) {

	mod := database.GetMod(name, dbType)

	if mod == nil {
		return nil, errors.New("mod does not exist")
	}

	fileDesc, err := download(mod.DownloadURL)

	if err != nil {
		return nil, errors.New("failed to download mod")
	}

	return fileDesc, nil
}

// Download a nonpacked mod zip file from url
func download(url string) (*os.File, error) {

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fileDesc, err := ioutil.TempFile("", "mod-*.zip")

	if err != nil {
		panic(err)
	}
	io.Copy(fileDesc, resp.Body)

	return fileDesc, nil
}
