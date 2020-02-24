package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func makeFile(folderPath string, fileName string) *os.File {
	fullPath := path.Join(folderPath, fileName)

	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	return file
}
func toJavaScriptClassName(fileName string) string {
	titleCase := strings.Title(fileName)
	return strings.ReplaceAll(titleCase, "-", "")
}

func makePluginScript(fileName string, file *os.File) {
	var scriptData string = fmt.Sprintf(`export default class %s extends Plugin {
	constructor(mod) {
		super();
		this.mod = mod;
	}

	async preload() {

	}

	async postload() {

	}

	async prestart() {

	}
}
`, toJavaScriptClassName(fileName))

	_, err := file.WriteString(scriptData)
	if err != nil {
		panic(err)
	}
}

// MakeScripts makes all scripts not found
func MakeScripts(folderPath string, pkg *Package) {
	if pkg.Preload != "" {
		preloadScript := makeFile(folderPath, pkg.Preload)
		preloadScript.Close()
	}

	if pkg.Postload != "" {
		postloadScript := makeFile(folderPath, pkg.Postload)
		postloadScript.Close()
	}

	if pkg.Prestart != "" {
		prestartScript := makeFile(folderPath, pkg.Prestart)
		prestartScript.Close()
	}

	if pkg.Plugin != "" {
		pluginScript := makeFile(folderPath, pkg.Plugin)
		stat, err := pluginScript.Stat()

		if err != nil {
			panic(err)
		} else if stat.Size() == 0 {
			makePluginScript(pkg.Name, pluginScript)
		}
		pluginScript.Close()
	}
}
