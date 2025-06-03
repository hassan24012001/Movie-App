package db_api

import (
	"context"

	"github.com/hassan/movie-db/object_defs"
)

type MoviesDb interface {
	GetTrendingMovies(ctx context.Context) ([]object_defs.Movie, error)
	GetUpcomingMovies(ctx context.Context) ([]object_defs.Movie, error)
	GetTopRatedMovies(ctx context.Context) ([]object_defs.Movie, error)
}
