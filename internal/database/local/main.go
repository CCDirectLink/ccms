package local

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/CCDirectLink/ccms/internal/database/generic"
	"github.com/CCDirectLink/ccms/internal/game"
)

// GetMods returns available mods
func GetMods() *generic.ModList {
	dir, _ := os.Getwd()
	dir, err := game.Find(dir)

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

		assetsPath, _ := game.Find(dir)

		modPath := path.Join(assetsPath, "mods", file.Name())
		if folderIsMod(modPath) {
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

func generateModEntryFrom(path string) *generic.ModEntry {

	_file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer _file.Close()

	var modData *generic.ModEntry
	err = json.NewDecoder(_file).Decode(&modData)

	if err != nil {
		panic(err)
	}

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
