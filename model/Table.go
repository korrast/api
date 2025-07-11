package model

import (
	"errors"

	uuid "github.com/google/uuid"
)

type Table struct {
	Id         uuid.UUID   `json:"id" gorm:"primaryKey; not null"`
	Title      string      `json:"title" gorm:"not null"`
	Columns    []Column    `gorm:"-"`
	Labels     []Label     `gorm:"-"`
	Milestones []Milestone `gorm:"-"`
}

func (t *Table) Init(title string) {
	t.Id = uuid.New()
	t.Title = title
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

func (t *Table) AddLabel(label Label) {
	t.Labels = append(t.Labels, label)
}

func (t *Table) AddMilestone(milestone Milestone) {
	t.Milestones = append(t.Milestones, milestone)
}
