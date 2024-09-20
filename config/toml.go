package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

func (c *Config) ReadFile(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	err = toml.NewDecoder(file).Decode(c)

	return err
}

func (c *Config) TOML() string {
	b, err := toml.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(b)
}
