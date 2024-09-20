package database

import (
	"fmt"

	"github.com/chlyNiklas/lou-taylor-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"db_name"`
	Port     int    `toml:"port"`
}
type DB struct {
	conn *gorm.DB
}

func New(cfg *Config) (db *DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	conn.AutoMigrate(&model.Event{})

	db = &DB{
		conn: conn,
	}

	return
}
