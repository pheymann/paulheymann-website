package movie

import "cdcexample/internal/inmemorydb"

type MovieService interface {
	ListAll() ([]Movie, error)
	Store(movie Movie) error
}

type ProductionMovieService struct {
	db inmemorydb.InMemoryDB[Movie]
}

func NewProductionMovieService(movies []Movie) *ProductionMovieService {
	// should be something like AWS DynamoDB but for the sake of simplicity we use
	// an in-memory DB
	db := inmemorydb.NewInMemoryDB(movies, func(movie Movie) string { return movie.ID })
	return &ProductionMovieService{db: db}
}
