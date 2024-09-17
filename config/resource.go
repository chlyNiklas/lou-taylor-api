package config

import (
	"time"

	"github.com/chlyNiklas/lou-taylor-api/authentication"
	"github.com/chlyNiklas/lou-taylor-api/image_service"
	"github.com/chlyNiklas/lou-taylor-api/model"
)

type Database struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

// Config holds all configuration values
type Config struct {
	Database *Database
	Security *authentication.Config
	Images   *image_service.Config
	BaseUrl  string
	SavePath string
}

func New() *Config {
	return &Config{

		Database: &Database{
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
		Security: &authentication.Config{
			Admin: &model.User{
				Name:     "admin",
				Password: "password",
			},
			JWTSecret:   []byte("my secret"),
			ValidPeriod: time.Hour * 3,
		},
		BaseUrl: "localhost:8080",
	}
}
