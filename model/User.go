package model

import (
	uuid "github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Username string
	Password string

	Tables []Table
}

func (u *User) Init(username string, password string) {
	u.Id = uuid.New()
	u.Username = username
	u.Password = password
}

func (u *User) AddTable(title string) {
	var t Table
	t.Init(title)
	u.Tables = append(u.Tables, t)
}
