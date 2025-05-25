package db

import (
	"fmt"

	"github.com/nix-united/golang-echo-boilerplate/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn(cfg)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open db connection: %w", err)
	}

	return db, nil
}

func dsn(c config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		c.Host, c.User, c.Password, c.Name, c.Port,
	)
}
