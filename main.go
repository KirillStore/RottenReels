package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"test_go/config"
	"test_go/controllers"
	"test_go/db"
	"test_go/middleware"
)

func main() {
	r := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Config loaded")

	_, err = db.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	moviesGroup := r.Group("/movies")
	{
		moviesGroup.GET("", controllers.GetAllMovies)
		moviesGroup.GET("/:id", controllers.GetMovieById)
	}

	authMoviesGroup := moviesGroup
	authMoviesGroup.Use(middleware.JWTAuthMiddleware())
	{
		authMoviesGroup.POST("", controllers.CreateMovie)
		authMoviesGroup.POST("/:id/reviews", controllers.CreateReview)
		//authMoviesGroup.GET("/:id/reviews", controllers.GetReviewsByMovieId)
		authMoviesGroup.POST("/:id/ratings", controllers.CreateRating)
		//authMoviesGroup.DELETE("/:id/", middleware.AdminMiddleware(), controllers.DeleteMovie)
		//authMoviesGroup.PUT("/:id/", middleware.AdminMiddleware(), controllers.UpdateMovie)
		//authMoviesGroup.DELETE("/:id/reviews", middleware.AdminMiddleware(), controllers.DeleteReview)

	}

	usersGroup := r.Group("/users")

	authUsersGroup := usersGroup
	authUsersGroup.Use(middleware.JWTAuthMiddleware())
	{
		authUsersGroup.GET("", controllers.GetAllUsers)
		authUsersGroup.GET("/:id", controllers.GetUserById)
		//authUsersGroup.DELETE("/:id", middleware.AdminMiddleware() controllers.DeleteUser)
		//authUsersGroup.PUT("/:id", middleware.AdminMiddleware(), controllers.UpdateUser) //в т.ч. смена роли юзера

	}

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controllers.LoginUser)
		authGroup.POST("/register", controllers.CreateUser)
	}

	r.Run(":8080")

}
