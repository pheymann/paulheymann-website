package listmovies

import (
	"cdcexample/internal/movie"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Hello struct {
	Message string
}

type Context struct {
	Movies movie.MovieService
}

func (ctx Context) Handle(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("listmovies.Handle")

	movies, err := ctx.Movies.ListAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movieJson, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(movieJson))
}
