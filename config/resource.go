package config

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
	JWTSecret []byte
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
		JWTSecret: []byte("my secret"),
	}
}
