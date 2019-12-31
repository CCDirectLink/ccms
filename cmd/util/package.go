package util

import (
	"encoding/json"
	"os"
	"path"
)

// Package represents a CrossCode Mod package.json file
type Package struct {

	// basics
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description,omitempty"`
	Module      bool   `json:"module,omitempty"`

	// script files
	Preload  string `json:"preload,omitempty"`
	Postload string `json:"postload,omitempty"`
	Prestart string `json:"prestart,omitempty"`
	Plugin   string `json:"plugin,omitempty"`

	// dependencies
	ModDep map[string]string `json:"ccmodDependencies,omitempty"`
	Dep    map[string]string `json:"dependencies,omitempty"`
}

// GetPackage takes a filePath and returns a Package or erro
func GetPackage(filePath string) (*Package, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	var data *Package
	err = json.NewDecoder(file).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// InitPackage creates an empty Package
func InitPackage() *Package {

	return &Package{
		Name:    "new-mod",
		Version: "0.0.0",
	}
}

// SavePackage takes a Package and saves it to the %folderPath%/package.json
func SavePackage(folderPath string, pkg *Package) (bool, error) {

	file, err := os.Create(path.Join(folderPath, "package.json"))
	if err != nil {
		return false, err
	}
	defer file.Close()

	pkgStr, _ := json.Marshal(pkg)
	file.WriteString(string(pkgStr))
	return true, nil
}
