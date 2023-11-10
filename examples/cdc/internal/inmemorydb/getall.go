package inmemorydb

func (db *InMemoryDB[E]) GetAll() []E {
	var result []E
	for _, v := range db.Entities {
		result = append(result, v)
	}
	return result
}
