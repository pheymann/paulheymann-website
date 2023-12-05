package inmemorydb

type InMemoryDB[E any] struct {
	Entities map[string]E
}

func NewInMemoryDB[E any](init []E, getId func(E) string) InMemoryDB[E] {
	db := InMemoryDB[E]{Entities: make(map[string]E)}

	for _, entity := range init {
		db.Entities[getId(entity)] = entity
	}

	return db
}
