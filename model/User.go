package model 

import (
  uuid "github.com/google/uuid"
)

type User struct {
  Id        uuid.UUID
  Pseudo    string
  Password  string
}

func (u *User) Init(pseudo string, password string) {
  u.Id = uuid.New()
  u.Pseudo = pseudo
  u.Password = password
}
