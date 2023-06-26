package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DBOpts struct {
	db    *gorm.DB
	redis *redis.Client
}

type RouterFunc func(router *gin.RouterGroup, dbOpts *DBOpts)

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

func SetupRoutes(db *gorm.DB, redis *redis.Client) *gin.Engine {
	router := gin.Default()
	router.GET("/check", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Server alive"}) })

	apiRouter := router.Group("api")
	for _, route := range registerRoutes() {
		routerGroup := apiRouter.Group(route.feature)
		route.registerRoute(routerGroup, &DBOpts{db, redis})
	}

	return router
}
