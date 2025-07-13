package database

import (
	"gorm.io/gorm"
)

func InsertLinkUserTable(db *gorm.DB, userId string, tableId string) error {
	return db.Table("users_tables").Create(map[string]interface{}{
		"userid":  userId,
		"tableid": tableId,
	}).Error
}

func InsertLinkTableColumn(db *gorm.DB, tableId string, columnId string) error {
	return db.Table("tables_columns").Create(map[string]interface{}{
		"tableid":  tableId,
		"columnid": columnId,
	}).Error
}
