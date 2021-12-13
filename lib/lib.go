package lib

import "os/exec"

func HasCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err == nil {
		return true
	} else {
		return false
	}
}

const (
	Ubuntu = 0
	CentOS = 1
)

func GetSystemName(os int) string {
	switch os {
	case Ubuntu:
		return "Ubuntu"
	case CentOS:
		return "CentOS"
	default:
		panic("No such system.")
	}
}

func GetSystemEnum() System {
	if HasCommand("apt-get") {
		return Ubuntu
	} else if HasCommand("yum") {
		return CentOS
	} else {
		panic("No such OS. Not supported.")
	}
}
