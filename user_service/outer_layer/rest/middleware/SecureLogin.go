package middleware

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	domainErrors "github.com/santaasus/errors-handler"
	security "shop/user_service/inner_layer/security"
)

func AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := os.ReadFile("config.json")
		if err != nil {
			appError := domainErrors.ThrowAppErrorWith(domainErrors.InternalServerError)
			_ = c.Error(appError)
			return
		}

		var config security.SecureConfig

		err = json.Unmarshal(file, &config)
		if err != nil {
			return
		}

		tokenString := c.GetHeader("Authorization")
		signature := []byte(config.Secure.JWTAcessSecure)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		var claims jwt.MapClaims
		_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return signature, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
