package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterItemRoutes(router *gin.RouterGroup, db *gorm.DB) {
	itemHandler := handlers.CreateItemHandler(db)

	router.Use(middlewares.ValidateSignatureMiddleware(), middlewares.AuthMiddleware())
	router.POST("", itemHandler.Create)
	router.GET("", itemHandler.Find)
	router.POST("transaction", itemHandler.Transaction)
	router.POST("movement", itemHandler.Movement)
}

var ItemRoute = Route{
	feature:       "items",
	registerRoute: RegisterItemRoutes,
}
