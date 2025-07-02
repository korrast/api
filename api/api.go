package api

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "main/stub"
  "main/model"
)

var datas stub.Stub

func InitApi(s stub.Stub) {
  router := gin.Default()

  if s != nil {
    datas = s
  }

  router = initializeRoutes(router)  

  router.Run()
}

func initializeRoutes(r *gin.Engine) *gin.Engine {
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
