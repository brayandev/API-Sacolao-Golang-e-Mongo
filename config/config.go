package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	_, err := toml.DecodeFile("config.toml", &c)

	if err != nil {
		log.Fatal(err)
	}
}
