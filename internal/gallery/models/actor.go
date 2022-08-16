package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Actor holds data related to an actor
// Note: we treat actor name as unique in this implementation
type Actor struct {
	gorm.Model
	Id   uint   `json:"-" gorm:"primary_key"`
	Name string `json:"actorName" gorm:"index:idx_name,unique"`
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (a *Actor) GoString() string {
	return fmt.Sprintf(`
{
	Name: %s,
}`,
		a.Name,
	)
}
