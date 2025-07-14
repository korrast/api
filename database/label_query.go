package database

import (
  "gorm.io/gorm"
)

func SelectLabelIdsFromTable(db *gorm.DB, tableID string) ([]string, error) {
  var labelIds []string

  if err := db.Table("tables_labels").Where("tableid = ?", tableID).Select("labelid").Find(&labelIds).Error; err != nil {
    return nil, err
  }

  return labelIds, nil
}
