package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	jwtHandler "github.com/santaasus/JWTToken-handler"
)

func ValidateJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		_, err := jwtHandler.VerifyTokenAndGetClaims(tokenString, jwtHandler.Access)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
