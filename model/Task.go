package model

import (
	uuid "github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Title       string
	Description string
}

func (t *Task) Init(title string, description string) {
	t.Id = uuid.New()
	t.Title = title
	t.Description = description
}
