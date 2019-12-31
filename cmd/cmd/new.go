package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/CCDirectLink/ccms/cmd/help"
	"github.com/CCDirectLink/ccms/cmd/util"
)

// New test
func New(wd string) {
	if len(os.Args) < 3 {
		help.New()
		return
	}

	name := strings.Join(os.Args[2:], " ")

	name = util.FormatPackageName(name)

	err := os.Mkdir(path.Join(wd, name), 0666)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	err = os.Mkdir(path.Join(wd, name, "assets"), 0666)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	fmt.Printf("Successfully created mod %s", name)
}
