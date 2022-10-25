package model

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	ID        uint   `json:"-" gorm:"primary_key"`
	Name      string `json:"name" gorm:"index:idx_name"`
	MovieDbId uint   `json:"-" gorm:"index:idx_movie_db_id"`
}

type ActorInput struct {
	Name string `json:"name"`
}

type Movie struct {
	gorm.Model
	ID          uint     `json:"-" gorm:"primary_key"`
	MovieID     string   `json:"movieId" gorm:"size:128;uniqueIndex"`
	Name        string   `json:"name"`
	Genre       string   `json:"genre"`
	ReleaseYear string   `json:"releaseYear"`
	Director    string   `json:"director"`
	Composer    string   `json:"composer"`
	Actors      []*Actor `json:"actors" gorm:"foreignKey:MovieDbId"`
}

type MovieInput struct {
	Name        string        `json:"name"`
	Genre       string        `json:"genre"`
	ReleaseYear string        `json:"releaseYear"`
	Director    string        `json:"director"`
	Composer    string        `json:"composer"`
	Actors      []*ActorInput `json:"actors"`
}
