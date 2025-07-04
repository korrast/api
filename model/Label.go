package model

import (
	uuid "github.com/google/uuid"
)

type Label struct {
	Id    uuid.UUID
	Title string
	Color string
}

func (l *Label) Init(title string, color string) {
	l.Id = uuid.New()
	l.Title = title
	l.Color = color
}
