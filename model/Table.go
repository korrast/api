package model

import (
	"errors"

	uuid "github.com/google/uuid"
)

type Table struct {
	Id      uuid.UUID
	Name    string
	Columns []Column
}

func (t *Table) Init(name string) {
	t.Id = uuid.New()
	t.Name = name
	var base [3]Column = BaseColumnsTemplate()
	t.Columns = append(t.Columns, base[0], base[1], base[2])
}

func (t *Table) AddColumn(name string) {
	var col Column
	// TODO
	// Change the color
	col.Init(name, "0xFFFFFF")
	t.Columns = append(t.Columns, col)
}

func (t *Table) AddTask(title string, description string, columnIndex int) error {
	if columnIndex >= len(t.Columns) {
		return errors.New("index out of bounds")
	}

	t.Columns[columnIndex].AddTask(title, description)
	return nil
}
