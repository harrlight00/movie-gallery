package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Movie holds data related to a movie
type Actor struct {
	gorm.Model
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (a *Actor) GoString() string {
	return fmt.Sprintf(`
{
    Id: %s,
	Name: %s,
}`,
		a.Id,
		a.Name,
	)
}
