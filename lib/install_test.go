package lib

import "testing"

func TestInstall(t *testing.T) {
	t.Skip("not available on DEV environment.")
	err := Install("hello")
	if err != nil {
		t.Error(err)
	}
}
