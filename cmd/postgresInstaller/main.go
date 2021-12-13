package main

import (
	"fmt"
	"installer/lib"
	"os"
)

func main() {
	err := lib.Install("postgresql postgresql-contrib")
	if err != nil {
		fmt.Errorf("Installer failed: %s", err.Error())
		os.Exit(1)
	}
}
