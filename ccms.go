package main

import (
	"fmt"
	"os"
	"path/filepath"

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

	basePackage, err := utils.GetPackage(filepath.Join(wd, "package.json"))

	hasPackage := true
	if err != nil {
		basePackage = utils.InitPackage()
		hasPackage = false
	}

	switch op {
	case "new", "n":
		wd = cmd.New(wd, basePackage)
		if wd != "" {
			utils.SavePackage(wd, basePackage)
		}
	case "init", "in":
		cmd.Init(basePackage)
		utils.SavePackage(wd, basePackage)
	case "install", "i":

		if !hasPackage {
			fmt.Println("main: could not find package.json in current directory")
			return
		}

		names := make([]string, 0)
		if flag.NArg() > 1 {
			fmt.Println("install: using cmd args to install")
			names = flag.Args()[1:]
		} else {
			fmt.Println("install: using package mod dep keys to install")
			for key := range basePackage.ModDep {
				names = append(names, key)
			}
		}

		for _, name := range names {
			stats := []*cmd.InstallStats{}
			stat := cmd.Install(wd, name, stats)

			if stat != nil {
				if stat.Err != nil {
					panic(stat.Err)
				}
			}

			if stat.Entry != nil {
				entry := stat.Entry
				basePackage.ModDep[entry.Name] = entry.Version
			}
		}
		utils.SavePackage(wd, basePackage)

	default:
		fmt.Printf("Invalid command: %s", op)
	}
}
