package mods

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/CCDirectLink/ccms/internal/database/dbtype"

	"github.com/CCDirectLink/ccms/internal/database"
)

// Download a nonpacked mod by name with db source
func Download(name string, db dbtype.DBType) (*os.File, error) {

	mod := database.GetMod(name, db)

	if mod == nil {
		return nil, errors.New("mod does not exist")
	}

	fileDesc, err := download(mod.Path)

	if err != nil {
		return nil, errors.New("failed to download mod")
	}

	return fileDesc, nil
}

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
