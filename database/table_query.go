package database

import (
	"errors"

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

func SelectTable(db *gorm.DB, userID string, tableID string) (*model.Table, error) {
	var tables []model.Table
	var tableIdInDb []string
	var columns, labels, milestones []string

	if err := db.Table("users_tables").Where("userid = ?", userID).Where("tableid = ?", tableID).Select("tableid").Find(&tableIdInDb).Error; err != nil {
		return nil, err
	}

	if len(tableIdInDb) == 0 {
		return nil, errors.New("no table with id" + tableID + "in db")
	}

	if err := db.Where("id = ?", tableIdInDb[0]).Find(&tables).Error; err != nil {
		return nil, err
	}

	if err := db.Table("tables_columns").Where("tableid = ?", tableID).Find(&columns).Error; err != nil {
		return nil, err
	}

	if err := db.Table("tables_labels").Where("tableid = ?", tableID).Find(&labels).Error; err != nil {
		return nil, err
	}

	if err := db.Table("tables_milestones").Where("tableid = ?", tableID).Find(&milestones).Error; err != nil {
		return nil, err
	}

	return &tables[0], nil
}
