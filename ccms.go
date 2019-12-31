package main

import (
	"fmt"
	"os"
	"path"

	"github.com/CCDirectLink/ccms/cmd/cmd"
	util "github.com/CCDirectlink/ccms/cmd/util"
)

func main() {

	if len(os.Args) == 1 {
		util.PrintHelp()
		return
	}

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	op := os.Args[1]

	basePackage, err := util.GetPackage(path.Join(wd, "package.json"))

	if err != nil {
		basePackage = util.InitPackage()
	}

	switch op {
	case "init":
		cmd.Init(basePackage)
		util.SavePackage(wd, basePackage)
	default:
		fmt.Printf("Invalid command: %s", op)
	}
}
