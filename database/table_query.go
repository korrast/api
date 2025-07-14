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

func SelectTable(db *gorm.DB, tableID string) (*model.Table, error) {
	var tables []model.Table
	var columns, labels, milestones []string

	if err := db.Where("id = ?", tableID).Find(&tables).Error; err != nil {
		return nil, err
	}

	if len(tables) == 0 {
		return nil, errors.New("no table with id" + tableID + "in db")
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

	//tables[0].Labels = labels
	//tables[0].Milestones = milestones
	//tables[0].Columns = columns

	return &tables[0], nil
}
