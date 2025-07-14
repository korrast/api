package database

import (
	"main/model"

	"gorm.io/gorm"
)

func InsertColumn(db *gorm.DB, column model.Column) error {
	return db.Create(&column).Error
}

func SelectColumnIdsFromTable(db *gorm.DB, tableID string) ([]string, error) {
  var columnIds []string

  if err := db.Table("tables_columns").Where("tableid = ?", tableID).Select("columnid").Find(&columnIds).Error; err != nil {
    return nil, err
  }

  return columnIds, nil
}
