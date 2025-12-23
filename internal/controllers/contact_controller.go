package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"github.com/yash-sojitra-20/address-book-backend/internal/services"
)

type ContactController struct {
	service *services.ContactService
}

func NewContactController(service *services.ContactService) *ContactController {
	return &ContactController{service}
}

func (c *ContactController) Create(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	var contact models.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := c.service.Create(userID, &contact); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, contact)
}

func (c *ContactController) GetAll(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	contacts, err := c.service.GetAll(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

func (c *ContactController) Update(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	id, _ := strconv.Atoi(ctx.Param("id"))

	var contact models.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := c.service.Update(userID, uint(id), &contact); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "contact updated"})
}

func (c *ContactController) Delete(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.Delete(userID, uint(id)); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "contact deleted"})
}