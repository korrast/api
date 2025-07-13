package database

import (
	"main/model"

	"gorm.io/gorm"
)

func InsertColumn(db *gorm.DB, column model.Column) error {
	return db.Create(&column).Error
}
