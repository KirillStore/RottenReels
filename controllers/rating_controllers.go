package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"test_go/db"
	"test_go/models"
	"test_go/utils"
)

func CreateRating(c *gin.Context) {
	var rating models.Rating

	// Логируем начало обработки запроса
	log.Println("Start processing CreateRating")

	// Привязка данных JSON к модели Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Rating data:", rating)

	// Получение данных пользователя из контекста
	claims, exists := c.Get("user")
	if !exists {
		log.Println("User claims not found")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userClaims := claims.(*utils.Claims)
	rating.UserID = uint(userClaims.UserID)
	log.Println("UserID from token:", rating.UserID)

	// Получение ID фильма из URL
	movieID := c.Param("id")
	if movieID == "" {
		log.Println("Movie ID is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "movie id is required"})
		return
	}
	movieIDint, err := strconv.Atoi(movieID)
	if err != nil {
		log.Println("Error converting movie ID to int:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rating.MovieID = uint(movieIDint)
	log.Println("MovieID:", rating.MovieID)

	// Попытка найти существующий рейтинг
	var existingRating models.Rating
	if err := db.DB.Where("user_id = ? AND movie_id = ?", rating.UserID, rating.MovieID).First(&existingRating).Error; err == nil {
		// Рейтинг существует, обновляем его
		existingRating.Score = rating.Score
		if err := db.DB.Save(&existingRating).Error; err != nil {
			log.Println("Error saving rating:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Println("Rating updated", existingRating)
		c.JSON(http.StatusOK, gin.H{"message": "Rating updated successfully", "rating": existingRating})
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Println("Error getting rating:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Сохранение нового рейтинга в базе данных
	rating.ID = 0 // Обнуляем ID, чтобы он генерировался автоматически
	if err := db.DB.Create(&rating).Error; err != nil {
		log.Println("Error during creation in DB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Rating created successfully")
	c.JSON(http.StatusCreated, gin.H{"message": "Rating created successfully", "rating": rating})
}
