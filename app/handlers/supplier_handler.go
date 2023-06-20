package handlers

import (
	"fmt"
	"net/http"
	"stockingapi/app/dtos"
	"stockingapi/app/models"
	"stockingapi/app/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SupplierHandler struct {
	validate           *validator.Validate
	supplierRepository *repositories.SupplierRepository
}

func CreateSupplierHandler(db *gorm.DB) *SupplierHandler {
	return &SupplierHandler{
		validate:           validator.New(),
		supplierRepository: repositories.CreateSupplierRepository(db),
	}
}

func (h *SupplierHandler) FindAll(c *gin.Context) {
	supplierModels, err := h.supplierRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, supplierModels)
}

func (h *SupplierHandler) Create(c *gin.Context) {
	var reqBody dtos.CreateSupplierRequestDTO

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplierExists, eSupplierExists := h.supplierRepository.FindOneSupplierByCode(reqBody.SupplierCode)
	if eSupplierExists != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": eSupplierExists.Error()})
		return
	}

	if supplierExists != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Supplier Code already taken."})
		return
	}

	supplierModel := models.SupplierModel{
		SupplierCode: reqBody.SupplierCode,
		SupplierName: reqBody.SupplierName,
	}

	if err := h.supplierRepository.Create(&supplierModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Supplier successfully created.", "id": supplierModel.ID})
}

func (h *SupplierHandler) Update(c *gin.Context) {
	var reqBody dtos.EditSupplierRequestDTO
	supplierCode := c.Param("code")

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplierExists, eSupplierExists := h.supplierRepository.FindOneSupplierByCode(supplierCode)

	if eSupplierExists != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": eSupplierExists.Error()})
		return
	}

	if supplierExists == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found."})
		return
	}

	editSupplierModel := models.EditSupplierModel{
		SupplierName: reqBody.SupplierName,
	}

	if err := h.supplierRepository.UpdateByCode(supplierCode, &editSupplierModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier successfully updated."})
}

func (h *SupplierHandler) Delete(c *gin.Context) {
	var supplierCode = c.Param("code")

	supplierExists, err := h.supplierRepository.FindOneSupplierByCode(supplierCode)
	fmt.Println("supplier exists", supplierExists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if supplierExists == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found."})
		return
	}

	if err := h.supplierRepository.DeleteByCode(supplierCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier successfully deleted."})
}
