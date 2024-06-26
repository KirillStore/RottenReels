package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "movie-rating-service/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    userController := controllers.NewUserController(db)
    r.POST("/register", userController.Register)
    r.POST("/login", userController.Login)
    r.POST("/logout", userController.Logout)

    movieController := controllers.NewMovieController(db)
    r.GET("/movies", movieController.GetAllMovies)
    r.POST("/movies", movieController.CreateMovie)
    r.GET("/movies/:id", movieController.GetMovieByID)
    r.PUT("/movies/:id", movieController.UpdateMovie)
    r.DELETE("/movies/:id", movieController.DeleteMovie)

    ratingController := controllers.NewRatingController(db)
    r.GET("/movies/:id/ratings", ratingController.GetAllRatingsForMovie)
    r.POST("/movies/:id/ratings", ratingController.CreateRating)
    r.GET("/ratings/:id", ratingController.GetRatingByID)
    r.PUT("/ratings/:id", ratingController.UpdateRating)
    r.DELETE("/ratings/:id", ratingController.DeleteRating)

    return r
}
