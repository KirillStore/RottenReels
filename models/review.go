package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Comment string `json:"comment"`
	Rating  uint   `json:"rating"` // Оценка от 1 до 10
	UserID  uint   `json:"user_id" binding:"required"`
	MovieID uint   `json:"movie_id" binding:"required"`
	User    User   `json:"user"`
}
