package database

import (
	"main/model"

	"gorm.io/gorm"
)

func AddUser(db *gorm.DB, user model.User) error {
	if err := db.Create(&user); err.Error != nil {
		return err.Error
	}

	return nil
}
