package database

import (
	"main/model"

	"gorm.io/gorm"
)

func InsertTable(db *gorm.DB, table model.Table) error {
	return db.Create(&table).Error
}
