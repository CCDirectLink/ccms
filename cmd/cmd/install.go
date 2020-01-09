package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/CCDirectLink/ccms/internal/database"
	"github.com/CCDirectLink/ccms/internal/database/dbtype"
	"github.com/CCDirectLink/ccms/internal/database/generic"
	"github.com/CCDirectLink/ccms/internal/mods"
	"github.com/CCDirectLink/ccms/internal/utils"
	"github.com/Masterminds/semver"
)

// InstallStats provides useful information for a specific
// mod installation
type InstallStats struct {
	Entry *generic.ModEntry
	Err   error
}

// Install a mod and add to package.json
func Install(wd string, name string, stats []*InstallStats) *InstallStats {

	fmt.Printf("Now installing %s\n", name)

	entry, err := install(wd, name)

	modStat := new(InstallStats)

	modStat.Entry = entry
	modStat.Err = err

	stats = append(stats, modStat)

	if err == nil {
		// go through dependencies and install
		packageData, err := utils.GetPackage(filepath.Join(modStat.Entry.Path, "package.json"))
		if err != nil && packageData.ModDep != nil {
			for depName := range packageData.ModDep {
				depStat := Install(wd, depName, stats)
				if depStat.Err != nil {
					panic(depStat.Err)
				}
			}
		}
	}

	return modStat
}

func installDependencies(gamePath string, name string) error {
	return nil
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
			fmt.Println("up to date")
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
