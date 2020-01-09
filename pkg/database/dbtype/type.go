package dbtype

// DBType to query
type DBType int

const (
	// LocalDB signfies local machine CrossCode mods
	LocalDB DBType = iota
	// CCModDB signifies Github Mods DB
	CCModDB
)

func (db DBType) String() string {
	return [...]string{"LocalDB", "CCModDB"}[db]
}
