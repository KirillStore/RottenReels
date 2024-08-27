package models

type Rating struct {
	Score   float64 `json:"score" binding:"required,gte=1,lte=5"`
	UserID  int     `json:"user_id"`
	MovieID int     `json:"movie_id"`
}
