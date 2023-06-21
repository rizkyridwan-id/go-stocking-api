package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterShelfRoutes(router *gin.RouterGroup, db *gorm.DB) {
	shelfHandler := handlers.CreateShelfHandler(db)

	router.Use(middlewares.ValidateSignatureMiddleware(), middlewares.AuthMiddleware())
	router.POST("", shelfHandler.Create)
	router.GET("", shelfHandler.FindAll)
	router.PUT(":code", shelfHandler.Update)
	router.DELETE(":code", shelfHandler.Delete)
}

var ShelfRoute = Route{
	feature:       "shelves",
	registerRoute: RegisterShelfRoutes,
}
