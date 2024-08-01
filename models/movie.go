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
	Reviews       []Review `json:"reviews" gorm:"foreignkey:MovieID"`
	AverageRating float64  `json:"average_rating" gorm:"-"`
}
