package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/CCDirectLink/ccms/pkg/logger"

	"github.com/CCDirectLink/ccms/pkg/database"
	"github.com/CCDirectLink/ccms/pkg/database/dbtype"
	"github.com/CCDirectLink/ccms/pkg/database/generic"
	"github.com/CCDirectLink/ccms/pkg/game"
	"github.com/CCDirectLink/ccms/pkg/mods"
	"github.com/CCDirectLink/ccms/pkg/utils"
	"github.com/Masterminds/semver"
)

// InstallStats provides useful information for a specific
// mod installation
type InstallStats struct {
	Entry *generic.ModEntry
	Err   string
}

// Install a mod and add to package.json
func Install(wd string, name string, stats []*InstallStats) *InstallStats {

	if shouldIgnoreDependency(name) {
		return &InstallStats{
			Entry: nil,
			Err:   "",
		}
	}

	gamePath, err := game.Find(wd)

	if err != nil {
		return &InstallStats{
			Entry: nil,
			Err:   logger.Critical("install", fmt.Sprintf("could not install %s. game directory not found", name)),
		}
	}

	logger.Log("install", fmt.Sprintf("now installing %s in \"%s\"", name, gamePath))

	entry := install(gamePath, name)

	modStat := new(InstallStats)

	modStat.Entry = entry

	stats = append(stats, modStat)

	if err == nil {
		// go through dependencies and install

		if entry == nil {
			modStat.Err = logger.Critical("install", fmt.Sprintf("%s somehow returns nil", name))
			return modStat
		}

		basePathInfo := fmt.Sprintf("base path is %s for %s", filepath.Join(entry.Path, "."), entry.Name)

		logger.Info("install", basePathInfo)

		packageData, err := utils.GetPackage(filepath.Join(entry.Path, "package.json"))

		if err != nil {
			errMsg := fmt.Sprintf("error installing %s: %s", entry.Name, err.Error())
			modStat.Err = logger.Critical("install", errMsg)
		}

		if packageData != nil && packageData.ModDep != nil {
			for depName := range packageData.ModDep {
				Install(wd, depName, stats)
			}
		}
	} else {
		modStat.Err = err.Error()
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

func install(gamePath string, name string) *generic.ModEntry {

	hasModLocally := database.HasMod(name, dbtype.LocalDB)

	// does it even exist?
	hasModGlobally := database.HasMod(name, dbtype.CCModDB)

	if hasModLocally {
		if hasModGlobally {
			localMod := database.GetMod(name, dbtype.LocalDB)
			globalMod := database.GetMod(name, dbtype.CCModDB)

			localVer, err := semver.NewVersion(localMod.Version)

			if err != nil {
				errMsg := fmt.Sprintf("encounter error: %s", err.Error())
				logger.Critical("install", errMsg)
				return nil
			}

			globalVer, err := semver.NewVersion(globalMod.Version)

			if globalVer.GreaterThan(localVer) {
				logger.Log("install", "updating...might break mods that depend on it")
			} else if globalVer.Equal(localVer) {
				logger.Log("install", "up to date")
				return localMod
			} else {
				logger.Critical("install", "downgrading is not support")
				return nil
			}
		} else {
			logger.Warn("install", fmt.Sprintf("mod %s only available locally", name))
			return database.GetMod(name, dbtype.LocalDB)
		}
	}
	// try downloading mods
	fileDesc, err := mods.Download(name, dbtype.CCModDB)

	if err != nil {
		errMsg := fmt.Sprintf("encounter error: %s", err.Error())
		logger.Critical("install", errMsg)
		return nil
	}

	filePath, err := mods.Extract(fileDesc)

	if err != nil {
		errMsg := fmt.Sprintf("encounter error: %s", err.Error())
		logger.Critical("install", errMsg)
		return nil
	}

	somePath := mods.FindPackage(filePath, name)

	if somePath == "" {
		errMsg := fmt.Sprintf("could not find package.json for %s", name)
		logger.Critical("install", errMsg)
		return nil
	}

	// copy

	err = mods.Copy(filepath.Dir(somePath), filepath.Join(gamePath, "mods", name, "."))

	if err != nil {
		errMsg := fmt.Sprintf("encounter error: %s", err.Error())
		logger.Critical("install", errMsg)
		return nil
	}
	return database.GetMod(name, dbtype.LocalDB)
}
