package mods

import (
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/CCDirectLink/ccms/pkg/utils"
)

// FindPackage by some basePath using the canonical name of the mod
// returns an empty string if not found
func FindPackage(basePath string, name string) string {
	libRegEx, e := regexp.Compile("package.json$")

	if e != nil {
		panic(e)
	}

	packagePath := ""

	e = filepath.Walk(basePath, func(walkPath string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == "node_modules" {
			return filepath.SkipDir
		}

		if err == nil && libRegEx.MatchString(info.Name()) {
			pkg, err := utils.GetPackage(walkPath)

			if err != nil {
				panic(err)
			}

			if pkg.Name == name {
				packagePath = walkPath

				return io.EOF
			}
		}
		return nil
	})

	if e == io.EOF {
		return packagePath
	}

	return ""
}
