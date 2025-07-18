package handler

import (
	"main/dto"
	"main/middleware"
	"main/response"
	"main/service"

	"github.com/gin-gonic/gin"
)

type TableHandler struct {
	tableService *service.TableService
}

func NewTableHandler(tableService *service.TableService) *TableHandler {
	return &TableHandler{
		tableService: tableService,
	}
}

func (h *TableHandler) CreateTable(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	var req dto.CreateTableRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, "Invalid request format: "+err.Error())
		return
	}

	table, err := h.tableService.CreateTable(userID, &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.CreatedResponse(c, "Table created successfully", table)
}

func (h *TableHandler) GetTables(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	tables, err := h.tableService.GetTables(userID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 200, "Tables fetched", tables)
}

func (h *TableHandler) GetTable(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	tableID := c.Param("id")

	table, err := h.tableService.GetTable(userID.String(), tableID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 200, "Table fetched", table)
}
