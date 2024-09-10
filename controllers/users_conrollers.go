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
	// Предзагрузка роли для каждого пользователя
	result := db.DB.Preload("Role").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	// Предзагрузка роли пользователя
	result := db.DB.Preload("Role").First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username and Password can't be empty"})
		return
	}

	if user.RoleID == 0 {
		var defaultRole models.Role
		if err := db.DB.Where("name = ?", "employee").First(&defaultRole).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to find default role"})
			return
		}
		user.RoleID = defaultRole.ID
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
		return
	}

	response := struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(201, response)
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
	result := db.DB.Where("username = ?", loginData.Username).Preload("Role").First(&user)
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

	// Генерация токена с учетом роли пользователя
	token, err := utils.GenerateJWT(user.Username, int(user.ID), user.Role.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := db.DB.First(&user, id).Preload("Role").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := db.DB.First(&user, id).Preload("Role").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated", "user": user})
}
