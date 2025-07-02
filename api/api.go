package api

import (
  "net/http"
  "log"
  "os"

  "github.com/gin-gonic/gin"

  "main/stub"
  "main/model"
)

var datas stub.Stub
var jwtSecret string

func InitApi() {
  router := gin.Default()

  if use_stub := os.Getenv("STUB_MODE"); use_stub == "true" {
    datas = &stub.StubImpl{}
    datas.CreateStubedData()  
  }

  if secret_token := os.Getenv("SECRET_TOKEN"); secret_token != "" {
    jwtSecret = secret_token
  } else {
    log.Fatalf("You have to specify a `SECRET_TOKEN` env variable somewhere")
  }

  router = initializeRoutes(router)  

  router.Run()
}

func initializeRoutes(r *gin.Engine) *gin.Engine {
    r.POST("/login", login)

    api := r.Group("/api")

    api.GET("/table", getTables)
    api.GET("/table/:id", getTable)
    api.GET("/table/:id/column", getColumns)

    return r
}

func getTables(c *gin.Context) {
  var data model.Table
  
  if datas != nil {
    data = datas.GetTable()
  }

  c.JSON(http.StatusOK, data)
}

func getTable(c *gin.Context) {
  c.JSON(http.StatusOK, "ok")
}

func getColumns(c *gin.Context) {
  c.JSON(http.StatusOK, "ok")
}
