package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/services"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, 400, "invalid json")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.Error(ctx, 400, err.Error())
		return
	}

	if err := c.authService.Register(req.Email, req.Password); err != nil {
		utils.Error(ctx, 400, err.Error())
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{"message": "user registered"})
	// ctx.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.Error(ctx, 400, err.Error())
		return
	}
	
	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		utils.Error(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(ctx, http.StatusOK, gin.H{
		"token": token,
	})
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })
}
