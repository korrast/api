package model

import (
	uuid "github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:" primaryKey; not null"`
	Username string    `json:"username" gorm:"not null"`
	Password string    `gorm:"not null"`

	Tables []Table `json:"tables" gorm:"-"`
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
