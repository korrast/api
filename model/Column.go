package model

import (
	uuid "github.com/google/uuid"
)

type Column struct {
	Id         uuid.UUID
	Name       string
	TaskNumber int
	Color      string
	Tasks      []Task
}

func (c *Column) Init(name string, color string) {
	c.Id = uuid.New()
	c.Name = name
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
	todo.Init("ToDo", "0xFFFFFF")
	inProgress.Init("In Progress", "0xFFFFFF")
	done.Init("Done", "0xFFFFFF")

	columns[0] = todo
	columns[1] = inProgress
	columns[2] = done

	return columns
}
