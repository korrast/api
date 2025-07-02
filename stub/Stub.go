package stub

import (
  "log"

  "main/model"
)

func CreateStubedData() model.Table {
  var table model.Table
  table.Init("Work Planer")
  if err := table.AddTask("MEP", "Finish production deployment", 1); err != nil {
    log.Fatalf("Error while loading stubed data " + err.Error())
  }

  return table
}
