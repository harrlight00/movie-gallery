package gallery

import (
	"github.com/google/uuid"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
)

// Helper class used to create a Movie object from a provided MovieInfo object
func createMovieFromMovieInfo(movieInfo models.MovieInfo) models.Movie {
	return models.Movie{MovieId: movieInfo.MovieId, Name: movieInfo.Name, Genre: movieInfo.Genre,
		ReleaseYear: movieInfo.ReleaseYear, Director: movieInfo.Director, Composer: movieInfo.Composer}
}

// Helper class used to create a MovieInfo object from a provided Movie object
func createMovieInfoFromMovie(movie models.Movie) models.MovieInfo {
	return models.MovieInfo{MovieId: movie.MovieId, Name: movie.Name, Genre: movie.Genre,
		ReleaseYear: movie.ReleaseYear, Director: movie.Director, Composer: movie.Composer}
}

// Helper class used to generate a UUID (for use when creating a Movie.MovieId)
func createUUID() string {
	return (uuid.New()).String()
}

// Helper class used to validate a UUID (for use when reading/updating)
func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
