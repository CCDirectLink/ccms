package local

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/CCDirectLink/ccms/internal/database/generic"
	"github.com/CCDirectLink/ccms/internal/game"
)

// GetMods returns available mods
func GetMods() *generic.ModList {

	var workdir *flag.Flag = flag.Lookup("wd")

	wd := workdir.Value.String()

	dir, err := game.Find(wd)

	if err != nil {
		panic(err)
	}

	modDir := game.GetModDirectory(dir)

	if modDir == "" {
		panic(errors.New("could not find mod directory"))
	}

	data := getMods(modDir)

	return data
}

func getMods(dir string) *generic.ModList {
	var list *generic.ModList

	// get all subdirectories
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	list = new(generic.ModList)

	list.Mods = make(map[string]generic.ModEntry)

	for _, file := range files {

		modPath := ""
		if file.Mode()&os.ModeSymlink != 0 {
			fullPath := filepath.Join(dir, file.Name())
			symbolPath, err := os.Readlink(fullPath)
			if err != nil {
				continue
			}
			modPath = symbolPath
		} else if file.IsDir() {
			modPath = path.Join(dir, file.Name())
		}

		if modPath != "" && folderIsMod(modPath) {
			packagePath := path.Join(modPath, "package.json")
			modEntry := generateModEntryFrom(packagePath)

			list.Mods[modEntry.Name] = *modEntry
		}

	}

	return list
}

func folderIsMod(dir string) bool {
	// check if it has a package.json
	// check if parent is mods

	packagePath := path.Join(dir, "package.json")

	if !fileExists(packagePath) {
		return false
	}

	return true
}

func generateModEntryFrom(packagePath string) *generic.ModEntry {

	_file, err := os.Open(packagePath)

	if err != nil {
		panic(err)
	}

	defer _file.Close()

	var modData *generic.ModEntry
	err = json.NewDecoder(_file).Decode(&modData)

	if err != nil {
		panic(err)
	}

	modData.Path = filepath.Dir(packagePath)

	return modData
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func folderExists(folderPath string) bool {
	stat, err := os.Stat(folderPath)

	if os.IsNotExist(err) {
		return false
	}

	if !stat.IsDir() {
		return false
	}
	return true
}
