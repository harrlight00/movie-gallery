package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Actor holds data related to an actor
type Actor struct {
	gorm.Model
	ActorId   string `json:"id"`
	Name      string `json:"actorName"`
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (a *Actor) GoString() string {
	return fmt.Sprintf(`
{
    ActorId: %s,
	Name: %s,
}`,
		a.ActorId,
		a.Name,
	)
}
