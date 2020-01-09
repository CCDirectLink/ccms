package ccmoddb

import (
	"encoding/json"
	"net/http"

	"github.com/CCDirectLink/ccms/pkg/database/generic"
)

var data *generic.ModList

var dbURL = "https://raw.githubusercontent.com/CCDirectLink/CCModDB/master/mods.json"

// GetMods returns available mods
func GetMods() *generic.ModList {
	if data != nil {
		return data
	}

	resp, err := http.Get(dbURL)

	if err != nil {
		return nil
	}

	var entryData *generic.ModList

	err = json.NewDecoder(resp.Body).Decode(&entryData)

	if err != nil {
		return nil
	}

	data = entryData
	return entryData
}
