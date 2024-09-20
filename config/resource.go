package config

import (
	"time"

	"github.com/chlyNiklas/lou-taylor-api/authentication"
	"github.com/chlyNiklas/lou-taylor-api/database"
	"github.com/chlyNiklas/lou-taylor-api/image_service"
	"github.com/chlyNiklas/lou-taylor-api/model"
)

// Config holds all configuration values
type Config struct {
	ConfigPath     string                 `toml:"-"`
	Database       *database.Config       `toml:"database" comment:"PostgreSQL connection"`
	Authentication *authentication.Config `toml:"security"`
	Images         *image_service.Config  `toml:"images" comment:"image endpoint"`
	BaseUrl        string                 `toml:"base_url"`
}

// Default returns a pointer to a default configuration
func Default() *Config {
	return &Config{
		ConfigPath: "./config.toml",
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
			ValidPeriod: time.Hour * 24,
		},
		BaseUrl: "localhost:8080",
	}
}
