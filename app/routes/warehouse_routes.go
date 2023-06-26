package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
)

func registerWarehouseRoutes(router *gin.RouterGroup, dbOpts *DBOpts) {
	warehouseHandler := handlers.CreateWarehouseHandler(dbOpts.db)

	router.Use(middlewares.ValidateSignatureMiddleware(), middlewares.AuthMiddleware())
	router.POST("", warehouseHandler.Create)
	router.GET("", warehouseHandler.FindAll)
	router.PUT(":code", warehouseHandler.Update)
	router.DELETE(":code", warehouseHandler.Delete)
}

var WarehouseRoute = Route{
	feature:       "warehouses",
	registerRoute: registerWarehouseRoutes,
}
