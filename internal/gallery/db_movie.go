package gallery;

import (
    models "github.com/harrlight00/movie-gallery/internal/gallery/models"
)

func getMovies(movie *models.Movie) error {
    var movies []models.Movie
    if result := db.Where("Name = ?", movie.Name).Find(&movies); result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

func insertMovie(movie *models.Movie) error {
	if result := db.Create(movie); result.Error != nil {
		return result.Error
	}
	return nil
}
