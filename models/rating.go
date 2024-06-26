package models

import "gorm.io/gorm"

type Rating struct {
    gorm.Model
    MovieID uint   `json:"movie_id"`
    UserID  uint   `json:"user_id"`
    Score   int    `json:"score"`
    Review  string `json:"review"`
}
