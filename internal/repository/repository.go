package repository

import (
	"errors"

	"gorm.io/gorm"
	"only-post-api/internal/auth"
	"only-post-api/internal/database"
)

func getUserByEmail(email string) (*auth.UserData, error) {
	var (
		db   = database.DB
		user auth.UserData
	)
	if err := db.Where(&auth.UserData{Email: email}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func getUserByUsername(username string) (*auth.UserData, error) {
	var (
		db   = database.DB
		user auth.UserData
	)
	if err := db.Where(&auth.UserData{Username: username}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
