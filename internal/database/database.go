package database

import (
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"only-post-api/config"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	port, err := strconv.ParseUint(config.Config("DB_PORT"), 10, 32)
	if err != nil {
		return err
	}

	cfg := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(cfg))
	if err != nil {
		return err
	}
	return nil
}
