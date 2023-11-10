package main

import (
	"cdcexample/internal/authorization"
	"cdcexample/internal/cmd/listmovies"
	"cdcexample/internal/cmd/storemovie"
	"cdcexample/internal/movie"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	router := mux.NewRouter()

	authorizer := authorization.NewProductionAuthorizer([]string{"abc", "def"})
	movies := movie.NewProductionMovieService([]movie.Movie{
		{ID: "1", Title: "The Godfather", ReleaseYear: 1972},
		{ID: "2", Title: "The Shawshank Redemption", ReleaseYear: 1994},
		{ID: "3", Title: "Schindler's List", ReleaseYear: 1993},
	})

	router.HandleFunc("/api/movie/all", listmovies.Context{Movies: movies}.Handle).Methods(http.MethodGet)
	router.HandleFunc("/api/movie", storemovie.Context{Authorizer: authorizer, Movies: movies}.Handle).Methods(http.MethodPut)

	log.Info().Msg("server up and runninng ...")
	log.Fatal().Err(http.ListenAndServe(":8080", router)).Msg("something went wrong")
	log.Info().Msg("server stopped")
}
