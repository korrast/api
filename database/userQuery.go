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

func GetUser(db *gorm.DB, username string, password string) ([]model.User, error) {
	var user []model.User
	result := db.Where("username = ?", username).Where("password = ?", password).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
