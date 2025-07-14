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

func (s *TableService) GetTables(userID uuid.UUID) (*[]dto.GetTablesResponse, error) {
	var res []dto.GetTablesResponse
	tables, err := database.SelectTables(s.db, userID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tables: %w", err)
	}

	for _, table := range tables {
		tableRes := dto.GetTablesResponse{
			Id:    table.Id.String(),
			Title: table.Title,
		}

		res = append(res, tableRes)
	}

	return &res, nil
}

func (s *TableService) GetTable(userID string, tableID string) (*dto.GetTableResponse, error) {
  var res dto.GetTableResponse

	table, err := database.SelectTable(s.db, userID, tableID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch table : %w", err)
	}

  columns, err := database.SelectColumnIdsFromTable(s.db, tableID)
  if err != nil {
		return nil, fmt.Errorf("failed to fetch column ids : %w", err)
  }

  labels, err := database.SelectLabelIdsFromTable(s.db, tableID)
  if err != nil {
		return nil, fmt.Errorf("failed to fetch column ids : %w", err)
  }  

  milestones, err := database.SelectMilestoneIdsFromTable(s.db, tableID)
  if err != nil {
		return nil, fmt.Errorf("failed to fetch column ids : %w", err)
  }  
  res.Id = table.Id.String()
  res.Title = table.Title
  res.ColumnIds = columns
  res.LabelIds = labels
  res.MilestoneIds = milestones

	return &res, nil
}
