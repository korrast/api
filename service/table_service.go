package service

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"main/database"
	"main/dto"
	"main/model"
)

type TableService struct {
	db *gorm.DB
}

func NewTableService(db *gorm.DB) *TableService {
	return &TableService{db: db}
}

func (s *TableService) CreateTable(userID uuid.UUID, req *dto.CreateTableRequest) (*model.Table, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var newTable model.Table
	newTable.Init(req.Title)

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := database.InsertTable(tx, newTable); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	if err := database.InsertLinkUserTable(tx, userID.String(), newTable.Id.String()); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to link table to user: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &newTable, nil
}

func (s *TableService) GetTables(userID uuid.UUID) (*[]model.Table, error) {
	tables, err := database.SelectTables(s.db, userID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tables: %w", err)
	}

	return &tables, nil
}
