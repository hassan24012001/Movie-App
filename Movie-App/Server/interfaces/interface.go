package interfaces

import "github.com/labstack/echo/v4"

type MovieService interface {
	TrendingMovies(ctx echo.Context) error
	UpcomingMovies(ctx echo.Context) error
	TopRatedMovies(ctx echo.Context) error
	SearchMovie(ctx echo.Context) error
	GetMovieDetail(ctx echo.Context) error
	GetMovieCredits(ctx echo.Context) error
	GetSimilarMovies(ctx echo.Context) error
	GetPersonDetails(ctx echo.Context) error
	GetPersonMovies(ctx echo.Context) error
}
