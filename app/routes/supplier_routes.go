package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterSupplierRoutes(router *gin.RouterGroup, dbOpts *DBOpts) {
	supplierHandler := handlers.CreateSupplierHandler(dbOpts.db)

	router.Use(middlewares.ValidateSignatureMiddleware(), middlewares.AuthMiddleware())
	router.POST("", supplierHandler.Create)
	router.GET("", supplierHandler.FindAll)
	router.PUT(":code", supplierHandler.Update)
	router.DELETE(":code", supplierHandler.Delete)
}

var SupplierRoute = Route{
	feature:       "suppliers",
	registerRoute: RegisterSupplierRoutes,
}
