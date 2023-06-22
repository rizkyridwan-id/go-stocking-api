package middlewares

import (
	"net/http"
	"stockingapi/app/helpers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		jwtHelper := helpers.CreateJWTHelper(15)

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]
		user, err := jwtHelper.ParseToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		c.Set("user", user.UserName)
		c.Next()
	}
}
