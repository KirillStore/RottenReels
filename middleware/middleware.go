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

		// Проверяем и удаляем префикс Bearer
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		// Валидируем токен
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Сохраняем информацию о пользователе в контексте
		c.Set("user", claims)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}
		userClaims, ok := claims.(*utils.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not contain an access token"})
			c.Abort()
			return
		}
		if userClaims.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not an admin"})
			c.Abort()
			return
		}
		c.Next()
	}
}
