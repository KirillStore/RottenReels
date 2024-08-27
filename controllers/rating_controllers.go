package controllers

import (
	"github.com/gin-gonic/gin"
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
	rating.UserID = int(userClaims.UserID)
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
	rating.MovieID = movieIDint
	log.Println("MovieID:", rating.MovieID)

	// Сохранение рейтинга в базе данных
	if err := db.DB.Create(&rating).Error; err != nil {
		log.Println("Error during creation in DB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Rating created successfully")
	c.JSON(http.StatusCreated, gin.H{"message": "Rating created successfully", "rating": rating})
}
