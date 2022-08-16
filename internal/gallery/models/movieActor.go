package models

import (
	"fmt"

	"gorm.io/gorm"
)

// MovieActor holds data related to an actors specific portrayal in a movie
type MovieActor struct {
	gorm.Model
	Id			  uint	   `json:"-" gorm:"primary_key"`
	Movie		  Movie	   `json:"movie" gorm:"foreignKey:MovieDbId;references:id"`
	MovieDbId	  uint     `json:"-" gorm:"index"`
	Actor		  Actor    `json:"actor" gorm:"foreignKey:ActorDbId;references:id"`
	ActorDbId	  uint     `json:"-" gorm:"index"`
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (ma *MovieActor) GoString() string {
    return fmt.Sprintf(`
{
    Movie: %s,
	Actor": %s,
}`,
        ma.Movie.GoString(),
		ma.Actor.GoString(),
    )
}
