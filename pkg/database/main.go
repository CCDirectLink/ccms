package database

import (
	"github.com/CCDirectLink/ccms/pkg/database/ccmoddb"
	"github.com/CCDirectLink/ccms/pkg/database/dbtype"
	"github.com/CCDirectLink/ccms/pkg/database/generic"
	"github.com/CCDirectLink/ccms/pkg/database/local"
)

// GetMods h
func GetMods(searchType dbtype.DBType) *generic.ModList {
	var mods *generic.ModList

	switch searchType {
	case dbtype.LocalDB:
		mods = local.GetMods()
	case dbtype.CCModDB:
		mods = ccmoddb.GetMods()
	default:
		mods = nil
	}
	return mods
}

// GetMod h
func GetMod(name string, searchType dbtype.DBType) *generic.ModEntry {
	mods := GetMods(searchType)

	if HasMod(name, searchType) {
		val, _ := mods.Mods[name]
		return &val
	}
	return nil
}

// HasMod h
func HasMod(name string, searchType dbtype.DBType) bool {

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
