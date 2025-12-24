package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"github.com/yash-sojitra-20/address-book-backend/internal/services"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
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
	// var contact CreateContactRequest
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	if err := utils.Validate.Struct(contact); err != nil {
		utils.Error(ctx, 400, err.Error())
		return
	}

	if err := c.service.Create(userID, &contact); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{
		"Data": contact,
	})
	// ctx.JSON(http.StatusCreated, contact)
}

func (c *ContactController) GetAll(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	contacts, err := c.service.GetAll(userID)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{
		"Data": contacts,
	})
	// ctx.JSON(http.StatusOK, contacts)
}

// Pagination
// func (c *ContactController) GetAll(ctx *gin.Context) {
// 	userID := ctx.GetUint("user_id")

// 	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
// 	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
// 	city := ctx.Query("city")

// 	contacts, err := c.service.GetPaginated(userID, page, limit, city)
// 	if err != nil {
// 		utils.Error(ctx, 500, "failed to fetch contacts")
// 		return
// 	}

// 	utils.Success(ctx, contacts)
// }

// Normal + Pagination 
// func (c *ContactController) GetAll(ctx *gin.Context) {
// 	userID := ctx.GetUint("user_id")

// 	// Read query params
// 	pageStr := ctx.Query("page")
// 	limitStr := ctx.Query("limit")
// 	city := ctx.Query("city")

// 	// If pagination params exist → use paginated method
// 	if pageStr != "" || limitStr != "" || city != "" {
// 		page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
// 		if err != nil {
// 			page = 1
// 		}

// 		limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
// 		if err != nil {
// 			limit = 10
// 		}

// 		contacts, err := c.service.GetPaginated(userID, page, limit, city)
// 		if err != nil {
// 			utils.Error(ctx, http.StatusInternalServerError, err.Error())
// 			return
// 		}

// 		utils.Success(ctx, http.StatusOK, gin.H{
// 			"page":  page,
// 			"limit": limit,
// 			"data":  contacts,
// 		})
// 		return
// 	}

// 	// Otherwise → fetch all contacts
// 	contacts, err := c.service.GetAll(userID)
// 	if err != nil {
// 		utils.Error(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	utils.Success(ctx, http.StatusOK, gin.H{
// 		"data": contacts,
// 	})
// }


func (c *ContactController) Update(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	id, _ := strconv.Atoi(ctx.Param("id"))

	var contact models.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	if err := c.service.Update(userID, uint(id), &contact); err != nil {
		utils.Error(ctx, http.StatusForbidden, err.Error())
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{"message": "contact updated"})
	// ctx.JSON(http.StatusOK, gin.H{"message": "contact updated"})
}

func (c *ContactController) Delete(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.Delete(userID, uint(id)); err != nil {
		utils.Error(ctx, http.StatusForbidden, err.Error())
		return
	}

	utils.Success(ctx, http.StatusOK, gin.H{"message": "contact deleted"})
	// ctx.JSON(http.StatusOK, gin.H{"message": "contact deleted"})
}