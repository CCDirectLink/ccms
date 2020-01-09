package generic

// ModEntry basic info
type ModEntry struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Path    string            `json:"archive_link,omitempty"`
	Hash    map[string]string `json:"hash,omitempty"`
}

// ModList is type representing a container of mods
type ModList struct {
	Mods map[string]ModEntry
}
