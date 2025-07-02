package main 

import (
  "fmt"

  "main/stub"
  "main/model"
)

func main() {
  var table model.Table = stub.CreateStubedData()

  fmt.Println("Table name : " + table.Name)
}
