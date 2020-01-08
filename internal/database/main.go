package database

import (
	"github.com/CCDirectLink/ccms/internal/database/ccmoddb"
	"github.com/CCDirectLink/ccms/internal/database/generic"
	"github.com/CCDirectLink/ccms/internal/database/local"
)

// GetMods h
func GetMods(searchType string) *generic.ModList {
	var mods *generic.ModList

	switch searchType {
	case "ccmoddb":
		mods = ccmoddb.GetMods()
	case "local":
		mods = local.GetMods()
	default:
		mods = nil
	}
	return mods
}

// GetMod h
func GetMod(name string, searchType string) *generic.ModEntry {
	mods := GetMods(searchType)

	if HasMod(name, searchType) {
		val, _ := mods.Mods[name]
		return &val
	}
	return nil
}

// HasMod h
func HasMod(name string, searchType string) bool {

	mods := GetMods(searchType)

	if mods == nil {
		return false
	}

	_, ok := mods.Mods[name]
	if ok {
		return true
	}

	return false
}
