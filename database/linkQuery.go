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
