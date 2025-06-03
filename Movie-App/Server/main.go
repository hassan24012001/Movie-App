package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hassan/movie-db/api_server/handlers"
	"github.com/hassan/movie-db/db_api"
	"github.com/hassan/movie-db/movies_client"
)

func main() {
	// migrate this to env
	db, err := sql.Open("postgres", "user=postgres dbname=movies password=tanzeem sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	e := echo.New()

	// Create an instance of Movies Wrapper
	moviesDb := db_api.NewMoviesWrapper(db)

	// set up movies client
	client := movies_client.NewMoviesClient()
	// Create an instance of the MovieServiceImpl struct.
	movieService := handlers.New(moviesDb, client)

	// Set up routes
	e.GET("/trending/movies", movieService.TrendingMovies)
	e.GET("/upcoming/movies", movieService.UpcomingMovies)
	e.GET("/top_rated/movies", movieService.TopRatedMovies)
	e.GET("/search/movie", movieService.SearchMovie)
	e.GET("/movie/:id", movieService.GetMovieDetail)
	e.GET("/movie/:id/credits", movieService.GetMovieCredits)
	e.GET("/movie/:id/similar", movieService.GetSimilarMovies)
	e.GET("/person/:id", movieService.GetPersonDetails)
	e.GET("/person/:id/credits", movieService.GetPersonMovies)

	e.Use(middleware.CORS())

	port := "0.0.0.0:8080"
	if err := e.Start(port); err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
