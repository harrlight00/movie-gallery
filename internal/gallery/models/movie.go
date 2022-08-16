package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Movie holds data related to a movie
type Movie struct {
	gorm.Model
	Id          uint         `json:"-" gorm:"primary_key"`
	MovieId     string       `json:"movieId" gorm:"size:128"`
	Name        string       `json:"name"`
	Genre       string       `json:"genre"`
	ReleaseYear string       `json:"releaseYear"`
	Director    string       `json:"director"`
	Composer    string       `json:"composer"`
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (m *Movie) GoString() string {
	return fmt.Sprintf(`
{
    MovieId: %s,
	Name: %s,
	Genre: %s,
    ReleaseYear: %s
    Director: %s,
    Composer: %s
}`,
		m.MovieId,
		m.Name,
		m.Genre,
		m.ReleaseYear,
		m.Director,
		m.Composer,
	)
}
