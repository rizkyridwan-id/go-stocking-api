package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterFunc func(router *gin.RouterGroup, db *gorm.DB)

type Route struct {
	feature       string
	registerRoute RouterFunc
}

// Register your new Feature Routes below
func registerRoutes() []*Route {
	return []*Route{
		&UserRoute,
		&WarehouseRoute,
		&ShelfRoute,
		&SupplierRoute,
		&GroupRoute,
		&ItemRoute,
	}
}

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/check", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Server alive"}) })

	apiRouter := router.Group("api")
	for _, route := range registerRoutes() {
		routerGroup := apiRouter.Group(route.feature)
		route.registerRoute(routerGroup, db)
	}

	return router
}
