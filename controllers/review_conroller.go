package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_go/db"
	"test_go/models"
)

// CreateReview allows an authorized user to leave a review.
func CreateReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&review)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}
