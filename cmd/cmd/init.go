package cmd

import (
	"fmt"

	"github.com/CCDirectLink/ccms/cmd/cli"
	"github.com/CCDirectLink/ccms/cmd/util"
)

func printFormat(formatStr string, str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf(formatStr, str)
}

// Init a mod package file
func Init(pkg *util.Package) {

	// basics
	name := cli.Prompt(fmt.Sprintf("name (%s):", pkg.Name))

	if name != "" {
		pkg.Name = util.FormatPackageName(name)
	}

	version := cli.Prompt(fmt.Sprintf("version (%s):", pkg.Version))

	if version != "" {
		pkg.Version = version
	}

	desc := cli.Prompt(printFormat("Old Description:%s\n", pkg.Description) + "Description:")

	if desc != "" {
		pkg.Description = desc
	}

	// script files

	preload := cli.Prompt("preload" + printFormat(" (%s)", pkg.Preload) + ":")

	if preload != "" {
		pkg.Preload = preload
	}

	postload := cli.Prompt("postload" + printFormat(" (%s)", pkg.Postload) + ":")

	if postload != "" {
		pkg.Postload = postload
	}

	prestart := cli.Prompt("prestart" + printFormat(" (%s)", pkg.Prestart) + ":")

	if prestart != "" {
		pkg.Prestart = prestart
	}

	plugin := cli.Prompt("plugin" + printFormat(" (%s)", pkg.Plugin) + ":")

	if plugin != "" {
		pkg.Plugin = plugin
	}

}
