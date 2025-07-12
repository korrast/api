package database

import (
	"main/model"

	"gorm.io/gorm"
)

func InsertTable(db *gorm.DB, table model.Table) error {
	return db.Create(&table).Error
}

func SelectTables(db *gorm.DB, userID string) ([]model.Table, error) {
	var tablesID []string
	var tables []model.Table

	if err := db.Table("users_tables").Where("userid = ?", userID).Select("tableid").Find(&tablesID).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id IN ?", tablesID).Find(&tables).Error; err != nil {
		return nil, err
	}

	return tables, nil
}
