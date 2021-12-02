package repository

import (
	"errors"
	"log"
	"strings"

	"gorm.io/gorm"
	"only-post-api/internal/auth"
	"only-post-api/internal/database"
)

func GetUserByEmail(email string) (*auth.UserData, error) {
	var (
		db   = database.DB
		user auth.UserData
	)
	if err := db.Table("users").Where(&auth.UserData{Email: email}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	user.Password = strings.TrimSpace(user.Password)
	return &user, nil
}

func GetUserByUsername(username string) (*auth.UserData, error) {
	var (
		db   = database.DB
		user auth.UserData
	)
	if err := db.Table("users").Where(&auth.UserData{Username: username}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	user.Password = strings.TrimSpace(user.Password)
	return &user, nil
}

func RegisterUser(user *auth.UserData) error {
	var (
		db = database.DB
	)
	if err := db.Table("users").Create(&user).Error; err != nil {
		return err
	}
	log.Printf("success register id: %d\n", user.ID)
	return nil
}
