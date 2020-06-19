package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database string
	Server   string
}

func (config *Config) Read() {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}
}
