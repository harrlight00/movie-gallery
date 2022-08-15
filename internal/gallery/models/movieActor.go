package models

import (
	"fmt"

	"gorm.io/gorm"
)

// MovieActor holds data related to an actors specific portrayal in a movie
type MovieActor struct {
	gorm.Model
	MovieId		  uint   `json:"movieId"`
	ActorId       string `json:"id"`
	CharacterName string `json:"characterName"`
    Actor         Actor  `json:"actor"`
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (ma *MovieActor) GoString() string {
	return fmt.Sprintf(`
{
    MovieId: %s,
    ActorId: %s,
	CharacterName: %s,
}`,
		ma.MovieId,
		ma.ActorId,
		ma.CharacterName,
	)
}
