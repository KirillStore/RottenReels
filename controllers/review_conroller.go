package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_go/db"
	"test_go/models"
)

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

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := db.DB.Where(review, id).First(&review).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "review deleted"})
}
