package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"main/database"
	"main/handler"
	"main/middleware"
	"main/service"
)

type Server struct {
	router        *gin.Engine
	dbManager     *database.Manager
	jwtSecret     string
	authHandler   *handler.AuthHandler
	tableHandler  *handler.TableHandler
	columnHandler *handler.ColumnHandler
}

func NewServer() (*Server, error) {
	dbManager, err := database.NewManager()
	if err != nil {
		return nil, err
	}

	jwtSecret := os.Getenv("SECRET_TOKEN")
	if jwtSecret == "" {
		log.Fatal("SECRET_TOKEN environment variable is required")
	}

	authService := service.NewAuthService(dbManager.GetDB(), jwtSecret)
	tableService := service.NewTableService(dbManager.GetDB())
	columnService := service.NewColumnService(dbManager.GetDB())

	authHandler := handler.NewAuthHandler(authService)
	tableHandler := handler.NewTableHandler(tableService)
	columnHandler := handler.NewColumnHandler(columnService)

	server := &Server{
		router:        gin.Default(),
		dbManager:     dbManager,
		jwtSecret:     jwtSecret,
		authHandler:   authHandler,
		tableHandler:  tableHandler,
		columnHandler: columnHandler,
	}

	server.setupRoutes()
	return server, nil
}

func (s *Server) setupRoutes() {
	s.router.POST("/register", s.authHandler.Register)
	s.router.POST("/login", s.authHandler.Login)

	api := s.router.Group("/api")
	api.Use(middleware.AuthMiddleware(s.jwtSecret))
	{
		api.POST("/tables", s.tableHandler.CreateTable)
		api.GET("/tables", s.tableHandler.GetTables)
		api.POST("/tables/:id/columns", s.columnHandler.CreateColumn)
	}
}

func (s *Server) Run(addr ...string) error {
	return s.router.Run(addr...)
}

func (s *Server) Close() error {
	return s.dbManager.Close()
}
