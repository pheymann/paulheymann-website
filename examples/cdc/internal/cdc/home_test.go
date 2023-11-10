package cdc

import (
	"cdcexample/internal/authorization"
	"cdcexample/internal/cmd/listmovies"
	"cdcexample/internal/movie"
	"net/http"
	"testing"
)

func Test_CDC_HomeNoMoviesFound(t *testing.T) {
	RunContracts(
		t,
		"/home/no_movies_found.yaml",
		runHomeContracts,
	)
}

func Test_CDC_HomeMultipleMoviesFound(t *testing.T) {
	RunContracts(
		t,
		"/home/multiple_movies_found.yaml",
		runHomeContracts,
	)
}

func runHomeContracts(
	r *http.Request,
	w http.ResponseWriter,
	movieService movie.MovieService,
	authorizer authorization.Authorizer,
) {
	ctx := listmovies.Context{Movies: movieService}

	ctx.Handle(w, r)
}
