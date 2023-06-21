package middlewares

import (
	"net/http"
	"stockingapi/app/helpers"

	"github.com/gin-gonic/gin"
)

func ValidateSignatureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.GetHeader("Signature")
		if signature == "" {
			signature = c.Query("signature")
		}

		timestamp := c.GetHeader("Timestamp")
		if timestamp == "" {
			timestamp = c.Query("timestamp")
		}

		isValidSignature := helpers.ValidateSignature(signature, timestamp, 60)

		if isValidSignature {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Signature."})
			c.Abort()
		}
	}
}
