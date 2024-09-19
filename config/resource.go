package config

import (
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"

	"github.com/chlyNiklas/lou-taylor-api/authentication"
	"github.com/chlyNiklas/lou-taylor-api/database"
	"github.com/chlyNiklas/lou-taylor-api/image_service"
	"github.com/chlyNiklas/lou-taylor-api/model"
)

// Config holds all configuration values
type Config struct {
	Database       *database.Config       `toml:"database" comment:"PostgreSQL connection"`
	Authentication *authentication.Config `toml:"security"`
	Images         *image_service.Config  `toml:"images" comment:"image endpoint"`
	BaseUrl        string                 `toml:"base_url"`
}

// Default returns a pointer to a default configuration
func Default() *Config {
	return &Config{

		Database: &database.Config{
			Host:     "localhost",
			User:     "username",
			Password: "password",
			Name:     "event_db",
			Port:     5432,
		},
		Images: &image_service.Config{
			Quality:  80,
			MaxWith:  2096,
			SavePath: ".tmp/",
		},
		Authentication: &authentication.Config{
			Admin: &model.User{
				Name:     "admin",
				Password: "password",
			},
			JWTSecret:   "my secret",
			ValidPeriod: time.Nanosecond,
		},
		BaseUrl: "localhost:8080",
	}
}

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
