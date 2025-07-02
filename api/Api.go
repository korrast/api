package api

import (
  "github.com/gin-gonic/gin"
)

func InitApi() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  router.Run()
}
