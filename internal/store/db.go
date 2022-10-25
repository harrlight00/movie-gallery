package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/harrlight00/movie-gallery/graph/model"
)

// Global used for accessing the DB
var DB *gorm.DB

func ConnectDB() {
	gormDb, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = gormDb

	if err := DB.AutoMigrate(&model.Movie{}, &model.Actor{}); err != nil {
		panic(err)
	}
}
