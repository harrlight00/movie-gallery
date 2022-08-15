package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)


// Movie holds data related to a movie
type Movie struct {
	gorm.Model
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"releaseDate"`
	Director    string    `json:"director"`
	Actors      []string  `json:"actors"`
	Composer    string    `json:"composer"`
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (m *Movie) GoString() string {
	actorListString := ""
	for i, actor := range m.Actors {
		actorListString += actor
		if i != (len(m.Actors) - 1) {
			actorListString += ", "
		}
	}

	return fmt.Sprintf(`
{
    Id: %s,
	Name: %s,
	Genre: %s,
    ReleaseDate: %s
    Director: %s,
    Actors: %s,
    Composer: %s
}`,
		m.Id,
		m.Name,
		m.Genre,
		m.ReleaseDate,
		m.Director,
		actorListString,
		m.Composer,
	)
}
