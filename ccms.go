package main

import (
	"fmt"
	"os"
	"path"

	"github.com/CCDirectLink/ccms/cmd/cmd"
	"github.com/CCDirectLink/ccms/cmd/help"
	"github.com/CCDirectLink/ccms/cmd/util"
	"github.com/CCDirectLink/ccms/internal/game"
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

	hasPackage := true
	if err != nil {
		basePackage = util.InitPackage()
		hasPackage = false
	}

	switch op {
	case "new":
		wd = cmd.New(wd, basePackage)
		if wd != "" {
			util.SavePackage(wd, basePackage)
		}
	case "init":
		cmd.Init(basePackage)
		util.SavePackage(wd, basePackage)
	case "install":

		if len(os.Args) < 3 {
			fmt.Println("main: must supply mod names")
			return
		}

		if !hasPackage {
			fmt.Println("main: could not find package.json in current directory")
			return
		}
		// first find game path
		gamePath, err := game.Find(wd)

		if err != nil {
			fmt.Println(err)
			return
		}

		names := os.Args[2:]

		// download
		err = cmd.Install(gamePath, names[0], basePackage)

		if err != nil {
			panic(err)
		}

		util.SavePackage(wd, basePackage)
	default:
		fmt.Printf("Invalid command: %s", op)
	}
}
