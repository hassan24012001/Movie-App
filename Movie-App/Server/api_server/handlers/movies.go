package handlers

import (
	"fmt"
	"net/http"

	"github.com/hassan/movie-db/db_api"
	"github.com/hassan/movie-db/interfaces"
	"github.com/hassan/movie-db/movies_client"

	"github.com/labstack/echo/v4"
)

// Add authz when refactoring and Add interceptors as well
// Inject deps when refactoring
type MovieServiceImpl struct {
	DataWrapper  *db_api.MoviesWrapper
	MoviesClient *movies_client.Client
}

func (m *MovieServiceImpl) TrendingMovies(ctx echo.Context) error {
	resp, err := m.DataWrapper.GetTrendingMovies(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Trending Movies Not Found")
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (m *MovieServiceImpl) UpcomingMovies(ctx echo.Context) error {
	resp, err := m.DataWrapper.GetUpcomingMovies(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Upcoming Movies Not Found")
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (m *MovieServiceImpl) TopRatedMovies(ctx echo.Context) error {
	resp, err := m.DataWrapper.GetTopRatedMovies(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Top Rated Movies Not Found")
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (m *MovieServiceImpl) SearchMovie(ctx echo.Context) error {
	query := ctx.QueryParam("query")

	params := make(map[string]interface{})
	params["include_adult"] = false
	params["language"] = "en-US"
	params["page"] = "1"
	params["query"] = query

	resp, err := m.MoviesClient.GetResponse("/search/movie", params)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Unable to get movies for the value")
	}

	respBytes, ok := resp.([]byte)
	if !ok {
		return ctx.String(http.StatusInternalServerError, "Invalid response type")
	}
	return ctx.Blob(http.StatusOK, "application/json", respBytes)
}

func (m *MovieServiceImpl) GetMovieDetail(ctx echo.Context) error {
	movieId := ctx.Param("id")

	resp, err := m.MoviesClient.GetResponse("/movie/"+movieId, nil)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Unable to get movie for the id %s", movieId))
	}

	respBytes, ok := resp.([]byte)
	if !ok {
		return ctx.String(http.StatusInternalServerError, "Invalid response type")
	}
	return ctx.Blob(http.StatusOK, "application/json", respBytes)
}

func (m *MovieServiceImpl) GetMovieCredits(ctx echo.Context) error {
	movieId := ctx.Param("id")

	resp, err := m.MoviesClient.GetResponse("/movie/"+movieId+"/credits", nil)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Unable to get movie credits for the movie id %s", movieId))
	}

	respBytes, ok := resp.([]byte)
	if !ok {
		return ctx.String(http.StatusInternalServerError, "Invalid response type")
	}
	return ctx.Blob(http.StatusOK, "application/json", respBytes)
}

func (m *MovieServiceImpl) GetSimilarMovies(ctx echo.Context) error {
	movieId := ctx.Param("id")

	resp, err := m.MoviesClient.GetResponse("/movie/"+movieId+"/similar", nil)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Unable to get similar movies for the movie id %s", movieId))
	}

	respBytes, ok := resp.([]byte)
	if !ok {
		return ctx.String(http.StatusInternalServerError, "Invalid response type")
	}
	return ctx.Blob(http.StatusOK, "application/json", respBytes)
}

func (m *MovieServiceImpl) GetPersonDetails(ctx echo.Context) error {
	personId := ctx.Param("id")

	resp, err := m.MoviesClient.GetResponse("/person/"+personId, nil)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Unable to get person details for id %s", personId))
	}

	respBytes, ok := resp.([]byte)
	if !ok {
		return ctx.String(http.StatusInternalServerError, "Invalid response type")
	}
	return ctx.Blob(http.StatusOK, "application/json", respBytes)
}

func (m *MovieServiceImpl) GetPersonMovies(ctx echo.Context) error {
	personId := ctx.Param("id")

	resp, err := m.MoviesClient.GetResponse("/person/"+personId+"/movie_credits", nil)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("Unable to get person movies for person id %s", personId))
	}

	respBytes, ok := resp.([]byte)
	if !ok {
		return ctx.String(http.StatusInternalServerError, "Invalid response type")
	}
	return ctx.Blob(http.StatusOK, "application/json", respBytes)
}

func New(dataWrapper *db_api.MoviesWrapper, moviesClient *movies_client.Client) *MovieServiceImpl {
	return &MovieServiceImpl{
		DataWrapper:  dataWrapper,
		MoviesClient: moviesClient,
	}
}

var _ interfaces.MovieService = (*MovieServiceImpl)(nil)
