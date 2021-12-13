package lib

import (
	"os/exec"
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
func Install(packName string) error {
	installTool := InstallCommand(
		System(GetSystemEnum()),
	)

	cmd := exec.Cmd{Args: []string{installTool, "install -y", packName}}
	err := cmd.Run()
	return err
}
