package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_go/db"
	"test_go/models"
	"test_go/utils"
)

func CreateRating(c *gin.Context) {
	var rating models.Rating

	// Привязка данных JSON к модели Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получение данных пользователя из контекста
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Приведение claims к нужному типу и получение UserID
	userClaims := claims.(*utils.Claims)
	rating.UserID = int(userClaims.UserID)

	// Получение ID фильма из URL
	movieID := c.Param("id")
	if movieID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Movie ID is required"})
		return
	}
	// Сохранение рейтинга в базе данных
	if err := db.DB.Create(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Rating created successfully", "rating": rating})
}
