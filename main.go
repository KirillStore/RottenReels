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

	authMoviesGroup := r.Group("/movies")
	authMoviesGroup.Use(middleware.JWTAuthMiddleware())
	{
		authMoviesGroup.POST("", controllers.CreateMovie)
		authMoviesGroup.POST("/:id/reviews", controllers.CreateReview)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.GET("", controllers.GetAllUsers)
		usersGroup.GET("/:id", controllers.GetUserById)
	}

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controllers.LoginUser)
		authGroup.POST("/register", controllers.CreateUser)
	}

	r.Run(":8080")

}
