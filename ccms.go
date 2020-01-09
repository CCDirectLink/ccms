package main

import (
	"fmt"
	"os"
	"path"

	"flag"

	"github.com/CCDirectLink/ccms/cmd/cmd"
	"github.com/CCDirectLink/ccms/cmd/help"
	"github.com/CCDirectLink/ccms/internal/utils"
)

func main() {

	var workdir *string
	dir, err := os.Getwd()

	workdir = flag.String("wd", dir, "a string")

	flag.Parse()

	if flag.NArg() < 1 {
		help.Default()
		return
	}

	wd := *workdir
	if err != nil {
		panic(err)
	}

	op := flag.Arg(0)

	basePackage, err := utils.GetPackage(path.Join(wd, "package.json"))

	hasPackage := true
	if err != nil {
		basePackage = utils.InitPackage()
		hasPackage = false
	}

	switch op {
	case "new":
		wd = cmd.New(wd, basePackage)
		if wd != "" {
			utils.SavePackage(wd, basePackage)
		}
	case "init":
		cmd.Init(basePackage)
		utils.SavePackage(wd, basePackage)
	case "install":

		if flag.NArg() < 2 {
			fmt.Println("main: must supply mod names")
			return
		}

		if !hasPackage {
			fmt.Println("main: could not find package.json in current directory")
			return
		}

		names := flag.Args()[1:]

		stats := []*cmd.InstallStats{}
		cmd.Install(wd, names[0], stats)

		if len(stats) > 0 {
			if stats[0].Err != nil {
				panic(stats[0].Err)
			}
		}

		/*if pkg.ModDep != nil {
			pkg.ModDep[entry.Name] = entry.Version
		}*/

		utils.SavePackage(wd, basePackage)
	default:
		fmt.Printf("Invalid command: %s", op)
	}
}
