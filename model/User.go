package model

import (
	uuid "github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Username string
	Password string
}

func (u *User) Init(username string, password string) {
	u.Id = uuid.New()
	u.Username = username
	u.Password = password
}
