package generic

// ModEntry basic info
type ModEntry struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	DownloadURL string            `json:"archive_link,omitempty"`
	Hash        map[string]string `json:"hash,omitempty"`
}

// ModList is type representing a local ModDB entry
type ModList struct {
	Mods map[string]ModEntry
}
