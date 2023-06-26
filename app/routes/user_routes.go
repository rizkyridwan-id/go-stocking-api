package routes

import (
	"stockingapi/app/handlers"
	"stockingapi/app/middlewares"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(router *gin.RouterGroup, dbOpts *DBOpts) {
	userHandler := handlers.CreateUserHandler(dbOpts.db, dbOpts.redis)
	router.Use(middlewares.ValidateSignatureMiddleware())
	router.POST("", userHandler.CreateUser)
	router.POST("/login", userHandler.LoginUser)
	router.POST("/refresh-token", userHandler.RefreshToken)

	router.Use(middlewares.AuthMiddleware())
	router.GET("", userHandler.GetUser)
	router.GET("/is-authorized", userHandler.IsAuthorized)
}

var UserRoute = Route{
	feature:       "users",
	registerRoute: registerUserRoutes,
}
