package cmd

import (
	"fmt"

	"github.com/CCDirectLink/ccms/internal/logger"
	"github.com/CCDirectLink/ccms/internal/utils"
)

// Uninstall a mod (or remove a reference to it from the corrent package)
func Uninstall(wd string, name string, pkg *utils.Package) {

	_, ok := pkg.ModDep[name]

	if !ok {
		logger.Warn("uninstall", fmt.Sprintf("mod %s did not depend on %s", pkg.Name, name))
		return
	}
	delete(pkg.ModDep, name)
	logger.Log("uninstall", fmt.Sprintf("dependency %s removed from %s", name, pkg.Name))
}
