package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"test_go/db"
	"test_go/models"
)

// GetAllMovies retrieves all movies with their average rating.
func GetAllMovies(c *gin.Context) {
	var movies []models.Movie
	result := db.DB.Preload("Reviews").Find(&movies)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Calculate average rating
	for i, movie := range movies {
		var totalRating int
		for _, review := range movie.Reviews {
			totalRating += review.Rating
		}
		if len(movie.Reviews) > 0 {
			movies[i].AverageRating = float64(totalRating) / float64(len(movie.Reviews))
		}
	}

	c.JSON(http.StatusOK, movies)
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

	result := db.DB.Create(&movie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}
