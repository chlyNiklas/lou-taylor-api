package config

import "github.com/chlyNiklas/lou-taylor-api/image_service"

// User holds all data of the admin user
type User struct {
	Name     string
	Password string
}

type Database struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

// Config holds all configuration values
type Config struct {
	Admin     *User
	Database  *Database
	Images    *image_service.ImageConfig
	JWTSecret []byte
	BaseUrl   string
	SavePath  string
}

func New() *Config {
	return &Config{
		Admin: &User{
			Name:     "admin",
			Password: "password",
		},
		Database: &Database{
			Host:     "localhost",
			User:     "username",
			Password: "password",
			Name:     "event_db",
			Port:     5432,
		},
		Images: &image_service.ImageConfig{
			Quality:  80,
			MaxWith:  2096,
			SavePath: ".tmp/",
		},
		JWTSecret: []byte("my secret"),
		BaseUrl:   "localhost:8080",
	}
}
