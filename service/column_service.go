package service

import (
	"fmt"

	"gorm.io/gorm"

	"main/database"
	"main/dto"
	"main/model"
)

type ColumnService struct {
	db *gorm.DB
}

func NewColumnService(db *gorm.DB) *ColumnService {
	return &ColumnService{db: db}
}

func (s *ColumnService) CreateColumn(tableId string, req *dto.CreateColumnRequest) (*model.Column, error) {
	var newColumn model.Column
	newColumn.Init(req.Title, req.Color)

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := database.InsertColumn(tx, newColumn); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create column: %w", err)
	}

	if err := database.InsertLinkTableColumn(tx, tableId, newColumn.Id.String()); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to link column to table: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &newColumn, nil
}
