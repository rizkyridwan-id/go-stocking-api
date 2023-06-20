package handlers

import (
	"fmt"
	"net/http"
	"stockingapi/app/dtos"
	"stockingapi/app/helpers"
	"stockingapi/app/models"
	"stockingapi/app/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type WarehouseHandler struct {
	validate            *validator.Validate
	warehouseRepository *repositories.WarehouseRepository
}

func CreateWarehouseHandler(db *gorm.DB) *WarehouseHandler {
	return &WarehouseHandler{
		validate:            validator.New(),
		warehouseRepository: repositories.CreateWarehouseRepository(db),
	}
}

func (h *WarehouseHandler) FindAll(c *gin.Context) {
	warehouseModels, err := h.warehouseRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, warehouseModels)
}

func (h *WarehouseHandler) Create(c *gin.Context) {
	var reqBody dtos.CreateWarehouseRequestDTO

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if testResult := helpers.RegexTestValidChar(reqBody.WarehouseCode); testResult == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Warehouse Code can only contains letters and numbers."})
		return
	}

	warehouseExists, _ := h.warehouseRepository.FindOneByWarehouseCode(reqBody.WarehouseCode)

	if warehouseExists != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Warehouse Code already taken."})
		return
	}

	warehouseModel := models.WarehouseModel{
		WarehouseCode: reqBody.WarehouseCode,
		WarehouseName: reqBody.WarehouseName,
	}

	errCreate := h.warehouseRepository.Create(&warehouseModel).Error
	if errCreate != nil {
		fmt.Printf("ERR: (WarehouseHandler)(Create) %s", errCreate.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Warehouse successfully created."})
}

func (h *WarehouseHandler) Update(c *gin.Context) {
	var reqBody dtos.EditWarehouseRequestDTO
	var warehouseCode = c.Param("code")

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	warehouseExists, _ := h.warehouseRepository.FindOneByWarehouseCode(warehouseCode)
	if warehouseExists == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found."})
		return
	}

	updateModel := models.EditWarehouseModel{
		WarehouseName: reqBody.WarehouseName,
	}
	err := h.warehouseRepository.UpdateByCode(warehouseCode, &updateModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Warehouse successfully updated.",
		"code":    warehouseCode,
	})
}

func (h *WarehouseHandler) Delete(c *gin.Context) {
	var warehouseCode = c.Param("code")

	warehouseExists, _ := h.warehouseRepository.FindOneByWarehouseCode(warehouseCode)
	if warehouseExists == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found."})
		return
	}

	err := h.warehouseRepository.DeleteByCode(warehouseCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Warehouse successfully deleted."})
}
