package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title         string   `gorm:"unique, not null" json:"title"`
	Description   string   `json:"desc"`
	Date          int      `json:"date"`
	Director      string   `json:"director"`
	AverageRating float64  `json:"average_rating" gorm:"-"`
	UserID        uint     `gorm:"index" json:"user_id"`
	Ratings       []Rating `gorm:"foreignKey:MovieID"`
	Reviews       []Review `gorm:"foreignKey:MovieID"`
}
