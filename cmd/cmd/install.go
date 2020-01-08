package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/CCDirectLink/ccms/internal/database/dbtype"
	"github.com/CCDirectLink/ccms/internal/database/generic"

	"github.com/CCDirectLink/ccms/cmd/util"
	"github.com/CCDirectLink/ccms/internal/database"
	"github.com/CCDirectLink/ccms/internal/mods"
	"github.com/Masterminds/semver"
)

// Install a mod and add to package.json
func Install(gamePath string, name string, pkg *util.Package) error {

	fmt.Println(pkg)

	entry, err := install(gamePath, name)

	if err != nil {
		return err
	}

	if pkg.ModDep == nil {
		pkg.ModDep = make(map[string]string)
	}

	pkg.ModDep[entry.Name] = entry.Version

	return nil
}

func installDependencies(gamePath string, name string) {

}

func install(gamePath string, name string) (*generic.ModEntry, error) {

	hasModLocally := database.HasMod(name, dbtype.LocalDB)

	// does it even exist?
	hasModGlobally := database.HasMod(name, dbtype.CCModDB)

	if !hasModGlobally {
		panic(errors.New("doesn't have mod"))
	}

	if hasModGlobally && hasModLocally {
		localMod := database.GetMod(name, dbtype.LocalDB)
		globalMod := database.GetMod(name, dbtype.CCModDB)

		localVer, err := semver.NewVersion(localMod.Version)

		if err != nil {
			return nil, err
		}

		globalVer, err := semver.NewVersion(globalMod.Version)

		if globalVer.GreaterThan(localVer) {
			fmt.Println("updating...might break mods that depend on it")
		} else if globalVer.Equal(localVer) {
			fmt.Println("update to date")
			return localMod, nil
		} else {
			fmt.Println("downgrading is not support")
			return nil, nil
		}
	}

	// try downloading mods
	fileDesc, err := mods.Download(name, dbtype.CCModDB)

	if err != nil {
		panic(err)
	}

	filePath, err := mods.Extract(fileDesc)

	if err != nil {
		panic(err)
	}

	somePath := mods.FindPackage(filePath, name)

	// copy

	err = mods.Copy(filepath.Dir(somePath), filepath.Join(gamePath, "mods", name, "."))

	if err != nil {
		panic(err)
	}
	return database.GetMod(name, dbtype.CCModDB), nil
}
