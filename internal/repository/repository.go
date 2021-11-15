package repository

import (
	"github.com/jinzhu/gorm"
	"only-post-api/internal/auth"
	"only-post-api/internal/database"
)

func getUserByEmail(email string) (*auth.UserData, error) {
	var (
		db   = database.DB
		user auth.UserData
	)
	if err := db.Where(&auth.UserData{Email: email}).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
