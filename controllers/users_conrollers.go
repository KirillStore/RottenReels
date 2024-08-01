package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"test_go/db"
	"test_go/models"
	"test_go/utils"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"message": result.Error.Error()})
	}
	c.JSON(200, users)
}
func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	result := db.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"message": "User not found"})
		} else {
			c.JSON(500, gin.H{"message": "Something went wrong"})
		}
		return
	}
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(400, gin.H{"message": "Username and Password can't be empty"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"message": result.Error.Error()})
	}
	c.JSON(201, user)
}

func LoginUser(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := db.DB.Where("username = ?", loginData.Username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
