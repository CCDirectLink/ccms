package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/CCDirectLink/ccms/internal/logger"

	"github.com/CCDirectLink/ccms/cmd/help"
	"github.com/CCDirectLink/ccms/internal/utils"
)

// New test
func New(wd string, pkg *utils.Package) string {
	if len(os.Args) < 3 {
		help.New()
		return ""
	}

	name := strings.Join(os.Args[2:], " ")

	name = utils.FormatPackageName(name)

	pkg.Name = name

	// https://chmod-calculator.com/

	newWd := path.Join(wd, name)
	err := os.Mkdir(newWd, 0755)

	if err != nil {
		if os.IsExist(err) {
			logger.Warn("new", fmt.Sprintf("mod %s already exists", name))
		} else {
			modFolderErr := fmt.Sprintf("failed to create mod folder...%s", err.Error())
			logger.Critical("new", modFolderErr)
			return ""
		}
	}

	err = os.Mkdir(path.Join(newWd, "assets"), 0755)

	if err != nil && !os.IsExist(err) {

		assetsFolderErr := fmt.Sprintf("failed to create assets folder...%s", err.Error())
		logger.Critical("new", assetsFolderErr)
		return ""
	}

	fmt.Printf("Successfully created mod %s", name)
	return newWd
}
