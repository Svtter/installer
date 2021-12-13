package lib

import (
	"testing"
)

func TestHasCommand(t *testing.T) {
	exist := HasCommand("ls")
	if !exist {
		t.Errorf("has command has error. ls not exist.")
	}

	exist = HasCommand("yum")
	if exist {
		t.Errorf("apt method should not exist.")
	}
}

func TestSystem(t *testing.T) {
	if GetSystemName(0) != "Ubuntu" {
		t.Errorf("Ubuntu is not defined.")
	}
}

func TestGetSystemEnum(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no such system. ")
		}
	}()
	GetSystemEnum()
}
