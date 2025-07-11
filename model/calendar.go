package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Calendar struct {
	Id     uuid.UUID
	Title  string
	Events []time.Time
}

func (c *Calendar) Init(title string) {
	c.Id = uuid.New()
	c.Title = title
}

func (c *Calendar) AddEvent(date time.Time) {
	c.Events = append(c.Events, date)
}
