package cdc

import (
	"cdcexample/internal/authorization"
	"cdcexample/internal/cmd/listmovies"
	"cdcexample/internal/cmd/storemovie"
	"cdcexample/internal/movie"
	"net/http"
	"strings"
	"testing"
)

func Test_CDC_CreateMovieUnauthorized(t *testing.T) {
	RunContracts(
		t,
		"/create_movie/unauthorized.yaml",
		runCreateHomeContract,
	)
}

func Test_CDC_CreateMovieCreated(t *testing.T) {
	RunContracts(
		t,
		"/create_movie/created.yaml",
		runCreateHomeContract,
	)
}

func runCreateHomeContract(
	r *http.Request,
	w http.ResponseWriter,
	movieService movie.MovieService,
	authorizer authorization.Authorizer,
) {
	storeCtx := storemovie.Context{
		Movies:     movieService,
		Authorizer: authorizer,
	}

	listCtx := listmovies.Context{
		Movies: movieService,
	}

	if strings.Contains(r.URL.Path, "all") {
		listCtx.Handle(w, r)
	} else {
		storeCtx.Handle(w, r)
	}
}
