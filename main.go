package main 

import (
  "fmt"
  "os"

  "main/api"
  "main/stub"
)

func main() {
  var s stub.Stub
  if use_stub := os.Getenv("STUB_MODE"); use_stub == "true" {
    s = &stub.StubImpl{}
    s.CreateStubedData()  
  }

  fmt.Println("Starting API")

  api.InitApi(s)
}
