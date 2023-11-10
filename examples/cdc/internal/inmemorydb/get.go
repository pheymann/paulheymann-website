package inmemorydb

func (db *InMemoryDB[E]) Get(id string) (E, bool) {
	v, ok := db.Entities[id]
	return v, ok
}
