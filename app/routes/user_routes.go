package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userHandler := handlers.CreateUserHandler(db)

	router.POST("", userHandler.CreateUser)
	router.POST("/login", userHandler.LoginUser)

	router.Use(middlewares.AuthMiddleware())
	router.GET("", userHandler.GetUser)
	router.GET("/is-authorized", userHandler.IsAuthorized)
}

var UserRoute = Route{
	feature:       "users",
	registerRoute: registerUserRoutes,
}
