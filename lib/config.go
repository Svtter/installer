package lib

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Name        string
	PackageList []string
	Username    string
	Password    string
}

func LoadConfig(name string) Config {
	var config Config
	f, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("read config fail.", err)
	}

	_, err = toml.Decode(string(f), &config)
	if err != nil {
		fmt.Println("Decode config error", err)
	}
	return config
}
