package handler

import (
	"main/dto"
	"main/response"
	"main/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, "Invalid request format: "+err.Error())
		return
	}

	if err := h.authService.Register(&req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.CreatedResponse(c, "User created successfully", nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, "Invalid request format: "+err.Error())
		return
	}

	authResponse, err := h.authService.Login(&req)
	if err != nil {
		response.UnauthorizedError(c, err.Error())
		return
	}

	response.OKResponse(c, "Login successful", authResponse)
}
