package auth

import (
	"github.com/jinzhu/gorm"
	"only-post-api/internal/database"
)

func getUserByEmail(email string) (*UserData, error) {
	var (
		db   = database.DB
		user UserData
	)
	if err := db.Where(&UserData{Email: email}).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
