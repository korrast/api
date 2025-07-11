package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"main/stub"
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
	r.POST("/register", register)

	api := r.Group("/api")

	api.Use(authMiddleware())

	api.POST("/tables", createTable)

	return r
}
