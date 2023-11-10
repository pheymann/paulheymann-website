package movie

func (s *ProductionMovieService) ListAll() ([]Movie, error) {
	return s.db.GetAll(), nil
}
