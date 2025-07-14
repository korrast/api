package database 

import (
  "gorm.io/gorm"
)

func SelectMilestoneIdsFromTable(db *gorm.DB, tableID string) ([]string, error) {
  var milestoneIds []string

  if err := db.Table("tables_milestones").Where("tableid = ?", tableID).Select("milestoneid").Find(&milestoneIds).Error; err != nil {
    return nil, err
  }

  return milestoneIds, nil
}
