package main

import (
	"fmt"
	"os"
	"path"

	"github.com/CCDirectLink/ccms/cmd/cmd"
	"github.com/CCDirectLink/ccms/cmd/help"
	"github.com/CCDirectLink/ccms/cmd/util"
)

func main() {

	if len(os.Args) == 1 {
		help.Default()
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
	case "new":
		cmd.New(wd)
	case "init":
		cmd.Init(basePackage)
		util.SavePackage(wd, basePackage)
	default:
		fmt.Printf("Invalid command: %s", op)
	}
}
