package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Movie holds data related to a movie in a format readable to the end user
type MovieInfo struct {
	gorm.Model
	MovieId     string    `json:"movieId"`
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
func (mi *MovieInfo) GoString() string {
	actorListString := ""
	for i, actor := range mi.Actors {
		actorListString += actor
		if i != (len(mi.Actors) - 1) {
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
		mi.MovieId,
		mi.Name,
		mi.Genre,
		mi.ReleaseDate,
		mi.Director,
		actorListString,
		mi.Composer,
	)
}
