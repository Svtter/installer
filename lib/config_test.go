package lib

import (
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(filepath.Dir(b))
)

func TestLoadConfig(t *testing.T) {
	config := LoadConfig(basePath + "/conf/postgres.toml")
	if config.Name != "postgres" {
		t.Errorf("parse error. %s", config.Name)
	}
}
