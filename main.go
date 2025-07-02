package main 

import (
  "fmt"

  "main/model"
)

func main() {
  var table model.Table
  table.Init("Table 1")

  fmt.Println("Id :" + table.Id.String() + ", Name :" + table.Name)
}
