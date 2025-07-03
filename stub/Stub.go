package stub

import (
	"log"

	"main/model"
)

type Stub interface {
	CreateStubedData()
	GetTable() model.Table
	GetUser() model.User
}

type StubImpl struct {
	Table model.Table
	User  model.User
}

func (s *StubImpl) CreateStubedData() {
	var table model.Table
	var user model.User
	table.Init("Work Planer")
	if err := table.AddTask("MEP", "Finish production deployment", 1); err != nil {
		log.Fatalf("Error while loading stubed data " + err.Error())
	}

	s.Table = table

	user.Init("Vincent", "123456789")

	s.User = user
}

func (s *StubImpl) GetTable() model.Table {
	return s.Table
}

func (s *StubImpl) GetUser() model.User {
	return s.User
}
