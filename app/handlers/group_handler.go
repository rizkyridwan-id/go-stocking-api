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

type GroupHandler struct {
	validate        *validator.Validate
	groupRepository *repositories.GroupRepository
}

func CreateGroupHandler(db *gorm.DB) *GroupHandler {
	return &GroupHandler{
		validate:        validator.New(),
		groupRepository: repositories.CreateGroupRepository(db),
	}
}

func (h *GroupHandler) FindAll(c *gin.Context) {
	groupModels, err := h.groupRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groupModels)
}

func (h *GroupHandler) Create(c *gin.Context) {
	var reqBody dtos.CreateGroupRequestDTO

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupExists, eGroupExists := h.groupRepository.FindOneByCode(reqBody.GroupCode)
	if eGroupExists != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": eGroupExists.Error()})
		return
	}

	if groupExists != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Group code already taken."})
		return
	}

	groupModel := models.GroupModel{
		GroupCode: reqBody.GroupCode,
		GroupName: reqBody.GroupName,
	}

	if err := h.groupRepository.Create(&groupModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Group successfully created.", "id": groupModel.ID})
}

func (h *GroupHandler) Delete(c *gin.Context) {
	var groupCode = c.Param("code")

	groupExists, err := h.groupRepository.FindOneByCode(groupCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if groupExists == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found."})
		return
	}

	if err := h.groupRepository.Delete(groupCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group successfully deleted."})
}
