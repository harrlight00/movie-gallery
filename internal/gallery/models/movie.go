package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Movie holds data related to a movie
type Movie struct {
	gorm.Model
	MovieId     string       `json:"movieId"`
	Name        string       `json:"name"`
	Genre       string       `json:"genre"`
	ReleaseDate time.Time    `json:"releaseDate"`
	Director    string       `json:"director"`
	MovieActors []MovieActor `json:"actors"`
	Composer    string       `json:"composer"`
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (m *Movie) GoString() string {
	return fmt.Sprintf(`
{
    Id: %s,
	Name: %s,
	Genre: %s,
    ReleaseDate: %s
    Director: %s,
    Composer: %s
}`,
		m.MovieId,
		m.Name,
		m.Genre,
		m.ReleaseDate,
		m.Director,
		m.Composer,
	)
}
