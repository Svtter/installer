// CLI installer.
package main

import (
	"fmt"
	"installer/lib"
	"os"
)

func main() {
	config := lib.LoadConfig("postgres")
	err := lib.Install(config.PackageList...)
	if err != nil {
		fmt.Errorf("Installer failed: %s", err.Error())
		os.Exit(1)
	}
}
