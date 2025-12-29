package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/config"
	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"github.com/yash-sojitra-20/address-book-backend/internal/services"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
)

type IAddressController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Export(ctx *gin.Context)
	ExportAsync(ctx *gin.Context)
	ExportCustom(ctx *gin.Context)
	GetFiltered(ctx *gin.Context)
}

type AddressController struct {
	service *services.AddressService
	cfg     *config.Config
}

func NewAddressController(service *services.AddressService, cfg *config.Config) *AddressController {
	return &AddressController{service, cfg}
}

func (c *AddressController) Create(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	var address models.Address
	// var address CreateAddressRequest
	if err := ctx.ShouldBindJSON(&address); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	if err := utils.Validate.Struct(address); err != nil {
		utils.Error(ctx, 400, err.Error())
		return
	}

	if err := c.service.Create(userID, &address); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{
		"Data": "address created",
	})

	// utils.Success(ctx, http.StatusCreated, gin.H{
	// 	"Data": address,
	// })
	// ctx.JSON(http.StatusCreated, address)
}

func (c *AddressController) GetAll(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	addresses, err := c.service.GetAll(userID)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{
		"Data": addresses,
	})
	// ctx.JSON(http.StatusOK, addresses)
}

func (c *AddressController) GetByID(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid address id")
		return
	}

	address, err := c.service.GetByID(userID, uint(id))
	if err != nil {
		utils.Error(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(ctx, http.StatusOK, gin.H{
		"data": address,
	})
}

func (c *AddressController) Update(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	id, _ := strconv.Atoi(ctx.Param("id"))

	var address models.Address
	if err := ctx.ShouldBindJSON(&address); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	if err := c.service.Update(userID, uint(id), &address); err != nil {
		utils.Error(ctx, http.StatusForbidden, err.Error())
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{"message": "address updated"})
	// ctx.JSON(http.StatusOK, gin.H{"message": "address updated"})
}

func (c *AddressController) Delete(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.Delete(userID, uint(id)); err != nil {
		utils.Error(ctx, http.StatusForbidden, err.Error())
		return
	}

	utils.Success(ctx, http.StatusOK, gin.H{"message": "address deleted"})
	// ctx.JSON(http.StatusOK, gin.H{"message": "address deleted"})
}

func (c *AddressController) Export(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	userEmail := ctx.GetString("user_email")

	if err := c.service.ExportAddresses(userID, userEmail, c.cfg); err != nil {
		utils.Error(ctx, 500, "failed to export addresses")
		return
	}

	utils.Success(ctx, http.StatusOK, "CSV sent to your email")
}

func (c *AddressController) ExportAsync(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	userEmail := ctx.GetString("user_email")

	// Start async export
	c.service.ExportAddressesAsync(userID, userEmail, c.cfg)

	// Respond immediately
	utils.Success(ctx, http.StatusAccepted, gin.H{
		"message": "Export started. CSV will be emailed shortly.",
	})
	// ctx.JSON(http.StatusAccepted, gin.H{
	// 	"success": true,
	// 	"message": "Export started. CSV will be emailed shortly.",
	// })
}

func (c *AddressController) ExportCustom(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	var req utils.CustomExportRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, 400, "invalid request payload")
		return
	}
	if err := utils.Validate.Struct(req); err != nil {
		utils.Error(ctx, 400, err.Error())
		return
	}

	c.service.ExportAddressesCustomAsync(userID, req.Fields, req.SendTo, c.cfg)

	utils.Success(ctx, 202, gin.H{
		"message": "Custom export started. CSV will be emailed shortly.",
	})
	// ctx.JSON(202, gin.H{
	// 	"message": "Custom export started. You will receive an email shortly.",
	// })
}

func (c *AddressController) GetFiltered(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	search := ctx.Query("search")
	city := ctx.Query("city")
	state := ctx.Query("state")
	country := ctx.Query("country")

	result, err := c.service.GetFilteredAddresses(
		int(userID), page, limit, search, city, state, country,
	)

	if err != nil {
		utils.Error(ctx, 500, "failed to fetch addresses")
		return
	}

	utils.Success(ctx, 200, result)
}
