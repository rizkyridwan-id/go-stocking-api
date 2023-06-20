package handlers

import (
	"net/http"
	"stockingapi/app/dtos"
	"stockingapi/app/models"
	"stockingapi/app/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ShelfHandler struct {
	validate            *validator.Validate
	shelfRepository     *repositories.ShelfRepository
	warehouseRepository *repositories.WarehouseRepository
}

func CreateShelfHandler(db *gorm.DB) *ShelfHandler {
	return &ShelfHandler{
		validate:            validator.New(),
		shelfRepository:     repositories.CreateShelfRepository(db),
		warehouseRepository: repositories.CreateWarehouseRepository(db),
	}
}

func (h *ShelfHandler) FindAll(c *gin.Context) {
	shelfModels, err := h.shelfRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelfModels)
}

func (h *ShelfHandler) Create(c *gin.Context) {
	var reqBody dtos.CreateShelfRequestDTO

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelfExists, errExists := h.shelfRepository.FindOneShelfByCode(reqBody.ShelfCode)
	if errExists != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errExists.Error()})
		return
	}

	if shelfExists != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Shelf Code already taken."})
		return
	}

	warehouseExists, errWarehouseExists := h.warehouseRepository.FindOneByWarehouseCode(reqBody.WarehouseCode)
	if errWarehouseExists != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errWarehouseExists})
		return
	}

	if warehouseExists == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Warehouse Code not found."})
		return
	}

	shelfModel := models.ShelfModel{
		ShelfCode: reqBody.ShelfCode,
		ShelfName: reqBody.ShelfName,
		Warehouse: *warehouseExists,
	}

	if err := h.shelfRepository.Create(&shelfModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Shelf successfully created.", "id": shelfModel.ID})
}

func (h *ShelfHandler) Update(c *gin.Context) {
	var reqBody dtos.EditShelfRequestDTO
	var shelfCode = c.Param("code")

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shelfExists, eShelfExists := h.shelfRepository.FindOneShelfByCode(shelfCode)

	if eShelfExists != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": eShelfExists.Error()})
		return
	}

	if shelfExists == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shelf not found."})
		return
	}

	editShelfModel := models.EditShelfModel{
		ShelfName: reqBody.ShelfName,
	}

	if err := h.shelfRepository.UpdateByCode(shelfCode, &editShelfModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shelf successfully updated."})
}

func (h *ShelfHandler) Delete(c *gin.Context) {
	var shelfCode = c.Param("code")

	shelfExist, _ := h.shelfRepository.FindOneShelfByCode(shelfCode)
	if shelfExist == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shelf not found."})
		return
	}

	if err := h.shelfRepository.DeleteByCode(shelfCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shelf successfully deleted."})
}
