package model

import (
	uuid "github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Title       string
	Description string
	Labels      []Label
	Milestone   Milestone
}

func (t *Task) Init(title string, description string) {
	t.Id = uuid.New()
	t.Title = title
	t.Description = description
}

func (t *Task) AddLabel(l Label) {
	t.Labels = append(t.Labels, l)
}

func (t *Task) SetMilestone(milestone Milestone) {
	t.Milestone = milestone
}
