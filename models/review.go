package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	RatingID int    `json:"rating_id"`
	Comment  string `json:"comment"`
	UserID   int    `json:"user_id" binding:"required"`
	MovieID  int    `json:"movie_id" binding:"required"`
}
