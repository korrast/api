package database

import (
	"main/model"

	"gorm.io/gorm"
)

func InsertUser(db *gorm.DB, user model.User) error {
	return db.Create(&user).Error
}

func GetUser(db *gorm.DB, username string, password string) ([]model.User, error) {
	var user []model.User
	result := db.Where("username = ?", username).Where("password = ?", password).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
