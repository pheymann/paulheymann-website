package movie

func (db ProductionMovieService) Store(movie Movie) error {
	db.db.Entities[movie.ID] = movie
	return nil
}
