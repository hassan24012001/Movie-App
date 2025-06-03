package db_api

import (
	"context"
	"database/sql"
	"log"

	"github.com/hassan/movie-db/object_defs"
	"github.com/lib/pq"

	"github.com/pkg/errors"
)

type MoviesWrapper struct {
	DB *sql.DB
}

func NewMoviesWrapper(db *sql.DB) *MoviesWrapper {
	return &MoviesWrapper{
		DB: db,
	}
}

func (w *MoviesWrapper) GetTrendingMovies(ctx context.Context) ([]object_defs.Movie, error) {
	rows, err := w.DB.Query("SELECT * FROM movies WHERE id BETWEEN 1 AND 20")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch trending movies")
	}
	defer rows.Close()

	var movies []object_defs.Movie
	var genreIDs pq.Int64Array

	for rows.Next() {
		var movie object_defs.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Adult,
			&movie.BackdropPath,
			&movie.MovieId,
			&movie.Title,
			&movie.OriginalLanguage,
			&movie.OriginalTitle,
			&movie.Overview,
			&movie.PosterPath,
			&movie.MediaType,
			&genreIDs,
			&movie.Popularity,
			&movie.ReleaseDate,
			&movie.Video,
			&movie.VoteAverage,
			&movie.VoteCount,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "error scanning rows")
		}

		movie.GenreIDs, err = scanIntArray(&genreIDs)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting GenreIDs")
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrapf(err, "error in movies rows")
	}

	log.Print("Get trending movies successfull")
	return movies, nil
}

func (w *MoviesWrapper) GetUpcomingMovies(ctx context.Context) ([]object_defs.Movie, error) {
	rows, err := w.DB.Query("SELECT * FROM movies WHERE id BETWEEN 21 AND 40")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch upcoming movies")
	}
	defer rows.Close()

	var movies []object_defs.Movie
	var genreIDs pq.Int64Array

	for rows.Next() {
		var movie object_defs.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Adult,
			&movie.BackdropPath,
			&movie.MovieId,
			&movie.Title,
			&movie.OriginalLanguage,
			&movie.OriginalTitle,
			&movie.Overview,
			&movie.PosterPath,
			&movie.MediaType,
			&genreIDs,
			&movie.Popularity,
			&movie.ReleaseDate,
			&movie.Video,
			&movie.VoteAverage,
			&movie.VoteCount,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "error scanning rows")
		}

		movie.GenreIDs, err = scanIntArray(&genreIDs)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting GenreIDs")
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrapf(err, "error in trending movies rows")
	}

	log.Print("Get upcoming movies successfull")
	return movies, nil
}

func (w *MoviesWrapper) GetTopRatedMovies(ctx context.Context) ([]object_defs.Movie, error) {
	rows, err := w.DB.Query("SELECT * FROM movies WHERE id BETWEEN 41 AND 60")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch top rated movies")
	}
	defer rows.Close()

	var movies []object_defs.Movie
	var genreIDs pq.Int64Array

	for rows.Next() {
		var movie object_defs.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Adult,
			&movie.BackdropPath,
			&movie.MovieId,
			&movie.Title,
			&movie.OriginalLanguage,
			&movie.OriginalTitle,
			&movie.Overview,
			&movie.PosterPath,
			&movie.MediaType,
			&genreIDs,
			&movie.Popularity,
			&movie.ReleaseDate,
			&movie.Video,
			&movie.VoteAverage,
			&movie.VoteCount,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "error scanning rows")
		}

		movie.GenreIDs, err = scanIntArray(&genreIDs)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting GenreIDs")
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrapf(err, "error in movies rows")
	}

	log.Print("Get top rated movies successfull")
	return movies, nil
}

var _ MoviesDb = (*MoviesWrapper)(nil)
