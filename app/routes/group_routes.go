package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterGroupRoutes(router *gin.RouterGroup, dbOpts *DBOpts) {
	groupHandler := handlers.CreateGroupHandler(dbOpts.db)

	router.Use(middlewares.ValidateSignatureMiddleware(), middlewares.AuthMiddleware())
	router.POST("", groupHandler.Create)
	router.GET("", groupHandler.FindAll)
	router.DELETE(":code", groupHandler.Delete)
}

var GroupRoute = Route{
	feature:       "groups",
	registerRoute: RegisterGroupRoutes,
}
