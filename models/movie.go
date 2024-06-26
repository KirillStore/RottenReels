package models

import "gorm.io/gorm"

type Movie struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    ReleaseDate string `json:"release_date"`
}
