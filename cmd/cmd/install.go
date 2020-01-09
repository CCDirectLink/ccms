package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/CCDirectLink/ccms/internal/database"
	"github.com/CCDirectLink/ccms/internal/database/dbtype"
	"github.com/CCDirectLink/ccms/internal/database/generic"
	"github.com/CCDirectLink/ccms/internal/game"
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

	if shouldIgnoreDependency(name) {
		return &InstallStats{
			Entry: nil,
			Err:   nil,
		}
	}

	gamePath, err := game.Find(wd)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Now installing %s in \"%s\"\n", name, gamePath)

	entry, err := install(gamePath, name)

	modStat := new(InstallStats)

	modStat.Entry = entry
	modStat.Err = err

	stats = append(stats, modStat)

	if err == nil {
		// go through dependencies and install

		if entry == nil {
			fmt.Println(entry)
			return modStat
		}

		fmt.Printf("base path is %s for %s\n", filepath.Join(entry.Path, "."), entry.Name)

		packageData, err := utils.GetPackage(filepath.Join(entry.Path, "package.json"))

		if err != nil {
			fmt.Printf("Error installing %s: %s", entry.Name, err)
		}

		if packageData != nil && packageData.ModDep != nil {
			for depName := range packageData.ModDep {
				Install(wd, depName, stats)
			}
		}
	}

	return modStat
}

func shouldIgnoreDependency(depName string) bool {
	ignore := [...]string{"ccloader", "crosscode"}

	for _, blacklist := range ignore {
		if blacklist == depName {
			return true
		}
	}
	return false
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

	if somePath == "" {
		return nil, fmt.Errorf("could not find package.json for %s", name)
	}

	// copy

	err = mods.Copy(filepath.Dir(somePath), filepath.Join(gamePath, "mods", name, "."))

	if err != nil {
		panic(err)
	}
	return database.GetMod(name, dbtype.LocalDB), nil
}
