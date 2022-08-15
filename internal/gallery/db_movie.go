package gallery

import (
	"errors"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"strings"
)

// Note: currently getMovies only accepts searching by one actor at a time
func getMovies(movieInfo *models.MovieInfo) ([]models.Movie, error) {
	var movies []models.Movie

	// Write query based on which fields are included
	query := `SELECT DISTINCT m.movieId, m.Name, m.Genre, m.ReleaseYear, m.Director, 
		m.Composer FROM movies m `

	fields := make([]string, 0)
	values := make([]interface{}, 0)

	if movieInfo.MovieId != "" {
		fields = append(fields, "m.movieId = ?")
		values = append(values, movieInfo.MovieId)
	}
	if movieInfo.Name != "" {
		fields = append(fields, "m.name = ?")
		values = append(values, movieInfo.Name)
	}
	if movieInfo.Genre != "" {
		fields = append(fields, "m.genre = ?")
		values = append(values, movieInfo.Genre)
	}
	if movieInfo.ReleaseYear != "" {
		fields = append(fields, "m.releaseYear = ?")
		values = append(values, movieInfo.ReleaseYear)
	}
	if movieInfo.Director != "" {
		fields = append(fields, "m.director = ?")
		values = append(values, movieInfo.Director)
	}
	if movieInfo.Composer != "" {
		fields = append(fields, "m.composer = ?")
		values = append(values, movieInfo.Composer)
	}
	if movieInfo.Actors != nil && len(movieInfo.Actors) != 0 {
		query += "LEFT JOIN movie_actors ma JOIN actors a "
		fields = append(fields, "a.name = ?")
		values = append(values, movieInfo.Actors[0])
	}

	if len(fields) == 0 {
		return nil, errors.New("Please specify a field to search against")
	}

	query += "WHERE " + strings.Join(fields, " AND ")

	if result := db.Exec(query, values...); result.Error != nil {
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
