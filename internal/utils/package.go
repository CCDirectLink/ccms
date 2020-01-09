package utils

import (
	"encoding/json"
	"os"
	"path"
	"regexp"
	"strings"
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

// GetPackage takes a filePath and returns a Package or error
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

	if data.ModDep == nil {
		data.ModDep = make(map[string]string)
	}

	return data, nil
}

// InitPackage creates an empty Package
func InitPackage() *Package {

	return &Package{
		Name:    "new-mod",
		Version: "0.0.0",
		ModDep:  make(map[string]string),
	}
}

// SavePackage takes a Package and saves it to the %folderPath%/package.json
func SavePackage(folderPath string, pkg *Package) (bool, error) {

	file, err := os.Create(path.Join(folderPath, "package.json"))
	if err != nil {
		return false, err
	}
	defer file.Close()

	pkgStr, _ := json.MarshalIndent(pkg, "", "\t")
	file.WriteString(string(pkgStr))
	return true, nil
}

// FormatPackageName to npm like spec
func FormatPackageName(name string) string {
	name = strings.ToLower(name)

	re := regexp.MustCompile(`/|\\`)
	name = re.ReplaceAllString(name, "")

	name = strings.Trim(name, "\t \n")

	re = regexp.MustCompile(`\s+`)
	name = re.ReplaceAllString(name, "-")

	return name
}
