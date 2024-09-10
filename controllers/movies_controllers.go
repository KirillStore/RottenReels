package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"test_go/db"
	"test_go/models"
	"test_go/utils"
)

// GetAllMovies retrieves all movies with their average rating.
func GetAllMovies(c *gin.Context) {
	var movies []models.Movie
	result := db.DB.Preload("Ratings").Find(&movies)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	for i, movie := range movies {
		var totalRating float64
		var countRating int
		var ratings []models.Rating

		resultRating := db.DB.Where("movie_id = ?", movie.ID).Find(&ratings)
		if resultRating.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": resultRating.Error.Error()})
			return
		}
		for _, rating := range ratings {
			totalRating += rating.Score
			countRating++
		}
		if countRating > 0 {
			movies[i].AverageRating = totalRating / float64(countRating)
		} else {
			movies[i].AverageRating = 0
		}
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})

}

// GetMovieById retrieves a movie by its ID.
func GetMovieById(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	result := db.DB.Preload("Reviews").First(&movie, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, movie)
}

// CreateMovie creates a new movie.
func CreateMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userClaims := claims.(*utils.Claims)
	movie.UserID = uint(userClaims.UserID) // Явное преобразование в uint, если требуется

	result := db.DB.Create(&movie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := db.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	if err := db.DB.Delete(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"movie": "Movie deleted successfully"})
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := db.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Save(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"movie": "Movie updated successfully"})
}
