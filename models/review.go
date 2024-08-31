package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Comment string `json:"comment"`
	UserID  uint   `json:"user_id" binding:"required"`
	MovieID uint   `json:"movie_id" binding:"required"`
}
