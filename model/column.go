package model

import (
	uuid "github.com/google/uuid"
)

type Column struct {
	Id         uuid.UUID `gorm:"primaryKey; not null"`
	Title      string    `gorm:"not null"`
	TaskNumber int       `gorm:"default: 0;column:tasknumber"`
	Color      string    `gorm:"not null;default: '#FFFFFF'"`
	Tasks      []Task    `gorm:"-"`
}

func (c *Column) Init(title string, color string) {
	c.Id = uuid.New()
	c.Title = title
	c.TaskNumber = 0
	c.Color = color
}

func (c *Column) AddTask(title string, description string) {
	var task Task
	task.Init(title, description)

	c.Tasks = append(c.Tasks, task)
}

func BaseColumnsTemplate() [3]Column {
	var columns [3]Column
	var todo, inProgress, done Column
	todo.Init("ToDo", "#00FF7F")
	inProgress.Init("In Progress", "#FFFACD")
	done.Init("Done", "#DDA0DD")

	columns[0] = todo
	columns[1] = inProgress
	columns[2] = done

	return columns
}
