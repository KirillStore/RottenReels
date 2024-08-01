package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test_go/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
			c.Abort()
			return
		}

		if strings.HasPrefix(tokenString, "Bearer") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer")
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
