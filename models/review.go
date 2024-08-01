package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating  int    `json:"rating" binding:"required, gte=1,lte=5"`
	Comment string `json:"comment"`
	UserID  int    `json:"user_id" binding:"required"`
	MovieID int    `json:"movie_id" binding:"required"`
}
