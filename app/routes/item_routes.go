package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(router *gin.RouterGroup, dbOpts *DBOpts) {
	itemHandler := handlers.CreateItemHandler(dbOpts.db)

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
