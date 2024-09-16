package database

import (
	"fmt"
	"github.com/chlyNiklas/lou-taylor-api/config"

	"github.com/chlyNiklas/lou-taylor-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

func New(cfg *config.Config) (db *DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	conn.AutoMigrate(&models.Event{})

	db = &DB{
		conn: conn,
	}

	return
}
