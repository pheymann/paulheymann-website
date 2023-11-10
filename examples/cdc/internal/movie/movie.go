package movie

type Movie struct {
	ID          string `json:"id" yaml:"id"`
	Title       string `json:"title" yaml:"title"`
	ReleaseYear int    `json:"releaseYear" yaml:"releaseYear"`
}
