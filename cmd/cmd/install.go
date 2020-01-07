package cmd

import (
	"fmt"
	"os"

	test "github.com/CCDirectLink/CCUpdaterCLI/cmd"
	"github.com/CCDirectLink/ccms/cmd/util"
)

// Install a mod
func Install(pkg *util.Package) {
	names := os.Args[2:]
	stats, err := test.Install(names)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*stats)

	fmt.Println(pkg)
}
