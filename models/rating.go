package models

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Score   float64 `json:"score" binding:"required,gte=1,lte=5"`
	UserID  uint    `json:"user_id"`
	MovieID uint    `json:"movie_id"`
}
