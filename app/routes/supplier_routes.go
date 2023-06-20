package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterSupplierRoutes(router *gin.RouterGroup, db *gorm.DB) {
	supplierHandler := handlers.CreateSupplierHandler(db)

	router.Use(middlewares.AuthMiddleware())
	router.POST("", supplierHandler.Create)
	router.GET("", supplierHandler.FindAll)
	router.PUT(":code", supplierHandler.Update)
	router.DELETE(":code", supplierHandler.Delete)
}

var SupplierRoute = Route{
	feature:       "suppliers",
	registerRoute: RegisterSupplierRoutes,
}
