package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerWarehouseRoutes(router *gin.RouterGroup, db *gorm.DB) {
	warehouseHandler := handlers.CreateWarehouseHandler(db)

	router.Use(middlewares.AuthMiddleware())
	router.POST("", warehouseHandler.Create)
	router.GET("", warehouseHandler.FindAll)
	router.PUT(":code", warehouseHandler.Update)
	router.DELETE(":code", warehouseHandler.Delete)
}

var WarehouseRoute = Route{
	feature:       "warehouses",
	registerRoute: registerWarehouseRoutes,
}
