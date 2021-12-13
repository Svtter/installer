package lib

import "testing"

func TestInstall(t *testing.T) {
	err := Install("hello")
	if err != nil {
		t.Error(err)
	}
}
