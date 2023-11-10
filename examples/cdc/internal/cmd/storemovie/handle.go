package storemovie

import (
	"cdcexample/internal/authorization"
	"cdcexample/internal/movie"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Context struct {
	Authorizer authorization.Authorizer
	Movies     movie.MovieService
}

func (ctx Context) Handle(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("storemovie.Handle")

	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	if !ctx.Authorizer.Authorize(bearerToken) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	var movie movie.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctx.Movies.Store(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
