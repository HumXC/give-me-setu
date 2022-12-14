package conf

import (
	"give-me-setu/util"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DataDir  string
	Library  string   `yaml:"library"`
	Database Database `yaml:"database"`
}

type Database struct {
	Driver   string `yaml:"driver"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func Get(path string) *Config {
	c := Config{
		Database: Database{
			Driver: "sqlite",
			Name:   "GiveMeSetu",
		},
	}

	if util.IsExist(path) {
		file, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(file, &c)
		if err != nil {
			panic(err)
		}
	}

	return &c
}
