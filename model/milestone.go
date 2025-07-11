package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Milestone struct {
	Id          uuid.UUID
	Title       string
	Description string
	EndDate     time.Time
}

func (m *Milestone) Init(title string, description string, endDate time.Time) {
	m.Id = uuid.New()
	m.Title = title
	m.Description = description
	m.EndDate = endDate
}
