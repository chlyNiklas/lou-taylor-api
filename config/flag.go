package config

import (
	"flag"
)

func (c *Config) ReadFlags() {
	flag.StringVar(&c.ConfigPath, "config", c.ConfigPath, "path of the toml configuration file")
	flag.StringVar(&c.BaseUrl, "base-ulr", c.BaseUrl, "base url of the api")

	// Database
	flag.StringVar(&c.Database.Host, "db-host", c.Database.Host, "host of the PostrgreSQL server")
	flag.IntVar(&c.Database.Port, "db-port", c.Database.Port, "port for the database connection")
	flag.StringVar(&c.Database.User, "db-user", c.Database.User, "username for database")
	flag.StringVar(&c.Database.Password, "db-password", c.Database.Password, "password for database")
	flag.StringVar(&c.Database.Name, "db-name", c.Database.Name, "name of the database")

	// Authentication Configuration
	flag.StringVar(&c.Authentication.Admin.Name, "admin-name", c.Authentication.Admin.Name, "administrator username")
	flag.StringVar(&c.Authentication.Admin.Password, "admin-password", c.Authentication.Admin.Password, "administrator password")
	flag.StringVar(&c.Authentication.JWTSecret, "jwt-secret", c.Authentication.JWTSecret, "JWT secret key for signing tokens")
	flag.DurationVar(&c.Authentication.ValidPeriod, "jwt-valid-period", c.Authentication.ValidPeriod, "valid period for JWT tokens")

	// Image Service Configuration
	flag.Float64Var(&c.Images.Quality, "image-quality", c.Images.Quality, "image quality (e.g., 80 for 80%)")
	flag.IntVar(&c.Images.MaxWith, "image-max-width", c.Images.MaxWith, "maximum width for uploaded images")
	flag.StringVar(&c.Images.SavePath, "image-save-path", c.Images.SavePath, "path to save uploaded images")

	flag.Parse()
}
