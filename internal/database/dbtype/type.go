package dbtype

type DBType int

const (
	LocalDB DBType = iota
	CCModDB
)

func (db DBType) String() string {
	return [...]string{"Local", "CCModDB"}[db]
}
