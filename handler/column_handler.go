package handler

import (
	"main/dto"
	"main/response"
	"main/service"

	"github.com/gin-gonic/gin"
)

type ColumnHandler struct {
	columnService *service.ColumnService
}

func NewColumnHandler(columnService *service.ColumnService) *ColumnHandler {
	return &ColumnHandler{
		columnService: columnService,
	}
}

func (h *ColumnHandler) CreateColumn(c *gin.Context) {
	var req dto.CreateColumnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, "Invalid request format: "+err.Error())
		return
	}

	column, err := h.columnService.CreateColumn(&req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.CreatedResponse(c, "Column created successfully", column)
}
