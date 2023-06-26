package handlers

import (
	"fmt"
	"net/http"
	"stockingapi/app/dtos"
	"stockingapi/app/helpers"
	"stockingapi/app/models"
	"stockingapi/app/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	Validate       *validator.Validate
	UserRepository *repositories.UserRepository
}

func CreateUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{UserRepository: repositories.CreateUserRepository(db), Validate: validator.New()}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	user, _ := c.MustGet("user").(helpers.Claims)
	fmt.Println(user)

	users, err := h.UserRepository.FindBy()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var reqBody dtos.RegisterRequestDTO
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if testResult := helpers.RegexTestValidChar(reqBody.UserId); testResult == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Id can only contains letters and numbers"})
		return
	}

	userExists, errorExists := h.UserRepository.FindOneByUserId(reqBody.UserId)
	if errorExists != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorExists.Error()})
		return
	}

	if userExists != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User Id already taken!"})
		return
	}

	fmt.Println("INFO: (UserHandler)(CreateUser) Find User Exists Result:", userExists)
	if userExists != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "UserId is already registered!"})
		return
	}

	passwordHashed, errEncrypt := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
	if errEncrypt != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errEncrypt.Error()})
		return
	}

	userModel := models.UserModel{
		UserId:   reqBody.UserId,
		UserName: reqBody.UserName,
		Password: string(passwordHashed),
		Level:    reqBody.Level,
	}

	err := h.UserRepository.Create(&userModel).Error
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully to create user"})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var reqBody dtos.LoginRequestDTO
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModel, _ := h.UserRepository.FindOneByUserId(reqBody.UserId)
	if userModel == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Data Not Found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(reqBody.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username/Password is wrong"})
		return
	}

	expiresIn := 3 * time.Hour
	jwtHelper := helpers.CreateJWTHelper(expiresIn)
	refreshTokenExpiresIn := 24 * time.Hour
	jwtHelperRefresh := helpers.CreateJWTHelper(refreshTokenExpiresIn)

	token := jwtHelper.GenerateToken(reqBody.UserId)
	tokenRefresh := jwtHelperRefresh.GenerateToken(reqBody.UserId)

	response := dtos.LoginResponseDTO{
		Message:               "Login Successfully",
		UserId:                userModel.UserId,
		Level:                 userModel.Level,
		Token:                 token,
		ExpiresIn:             uint(expiresIn.Seconds()),
		RefreshToken:          tokenRefresh,
		RefreshTokenExpiresIn: uint(refreshTokenExpiresIn.Seconds()),
	}

	c.JSON(http.StatusOK, response)
	return
}

// Example of extracting user after authorization
func (h *UserHandler) IsAuthorized(c *gin.Context) {
	user := c.MustGet("user").(*helpers.Claims)
	c.JSON(http.StatusOK, gin.H{"message": "Hello ," + user.UserName})
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var reqBody dtos.RefreshTokenRequestDTO
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Validate.Struct(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expiresIn := 3 * time.Hour
	jwtHelper := helpers.CreateJWTHelper(expiresIn)

	user, err := jwtHelper.ParseToken(reqBody.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Refresh Token."})
		return
	}

	newToken := jwtHelper.GenerateToken(user.UserName)

	refreshTokenExpiresIn := 24 * time.Hour
	jwtHelperRefresh := helpers.CreateJWTHelper(refreshTokenExpiresIn)

	newRefreshToken := jwtHelperRefresh.GenerateToken(user.UserName)

	c.JSON(http.StatusOK, dtos.RefreshTokenResponseDTO{
		Token:                 newToken,
		ExpiresIn:             uint(expiresIn.Seconds()),
		RefreshToken:          newRefreshToken,
		RefreshTokenExpiresIn: uint(refreshTokenExpiresIn.Seconds()),
	})
}
