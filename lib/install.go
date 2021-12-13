package lib

import (
	"os/exec"
	"strings"
)

type System int

// InstallCommand get install command.
func InstallCommand(system System) string {
	switch system {
	case Ubuntu:
		return "apt-get"
	case CentOS:
		return "yum"
	default:
		panic("UNKNOWN system")
	}
}

// Install install package
func Install(packNames ...string) error {
	installTool := InstallCommand(
		System(GetSystemEnum()),
	)

	packName := strings.Join(packNames, " ")
	cmd := exec.Cmd{Args: []string{installTool, "install -y", packName}}
	err := cmd.Run()
	return err
}
