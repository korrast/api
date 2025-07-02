package stub

import (
  "log"

  "main/model"
)

type Stub interface {
  CreateStubedData()
  GetTable() model.Table
}

type StubImpl struct {
  Table model.Table 
}

func (s *StubImpl) CreateStubedData() {
  var table model.Table
  table.Init("Work Planer")
  if err := table.AddTask("MEP", "Finish production deployment", 1); err != nil {
    log.Fatalf("Error while loading stubed data " + err.Error())
  }

  s.Table = table 
}

func (s *StubImpl) GetTable() model.Table {
  return s.Table
}
