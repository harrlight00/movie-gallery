package gallery

import (
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func testCreateMovieFromMovieInfo(t *testing.T) {
	assert := assert.New(t)

	testMovie := models.Movie{
		Id:          1,
		MovieId:     "testMovieId",
		Name:        "Tenet",
		Genre:       "Action",
		ReleaseYear: "2020",
		Director:    "Christopher Nolan",
		Composer:    "Ludwig Goransson",
	}
	testMovieInfo := models.MovieInfo{
		Id:          1,
		MovieId:     "testMovieId",
		Name:        "Tenet",
		Genre:       "Action",
		ReleaseYear: "2020",
		Director:    "Christopher Nolan",
		Composer:    "Ludwig Goransson",
		Actors:      make([]string, 0),
	}

	actualMovie := createMovieFromMovieInfo(testMovieInfo)
	assert.Equal(testMovie.Id, actualMovie.Id, "Id is equivalent")
	assert.Equal(testMovie.MovieId, actualMovie.MovieId, "MovieId is equivalent")
	assert.Equal(testMovie.Name, actualMovie.Name, "Name is equivalent")
	assert.Equal(testMovie.Genre, actualMovie.Genre, "Genre is equivalent")
	assert.Equal(testMovie.ReleaseYear, actualMovie.ReleaseYear, "ReleaseYear is equivalent")
	assert.Equal(testMovie.Director, actualMovie.Director, "Director is equivalent")
	assert.Equal(testMovie.Composer, actualMovie.Composer, "Composer is equivalent")
}

func testCreateMovieInfoFromMovie(t *testing.T) {
	assert := assert.New(t)

	testMovie := models.Movie{
		Id:          1,
		MovieId:     "testMovieId",
		Name:        "Tenet",
		Genre:       "Action",
		ReleaseYear: "2020",
		Director:    "Christopher Nolan",
		Composer:    "Ludwig Goransson",
	}
	testMovieInfo := models.MovieInfo{
		Id:          1,
		MovieId:     "testMovieId",
		Name:        "Tenet",
		Genre:       "Action",
		ReleaseYear: "2020",
		Director:    "Christopher Nolan",
		Composer:    "Ludwig Goransson",
		Actors:      make([]string, 0),
	}

	actualMovieInfo := createMovieInfoFromMovie(testMovie)
	assert.Equal(testMovieInfo.Id, actualMovieInfo.Id, "Id is equivalent")
	assert.Equal(testMovieInfo.MovieId, actualMovieInfo.MovieId, "MovieId is equivalent")
	assert.Equal(testMovieInfo.Name, actualMovieInfo.Name, "Name is equivalent")
	assert.Equal(testMovieInfo.Genre, actualMovieInfo.Genre, "Genre is equivalent")
	assert.Equal(testMovieInfo.ReleaseYear, actualMovieInfo.ReleaseYear, "ReleaseYear is equivalent")
	assert.Equal(testMovieInfo.Director, actualMovieInfo.Director, "Director is equivalent")
	assert.Equal(testMovieInfo.Composer, actualMovieInfo.Composer, "Composer is equivalent")
}

func testCreateUUID(t *testing.T) {
	assert := assert.New(t)
	uuid := createUUID()

	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	assert.True(r.MatchString(uuid), "UUID generated is of valid format")
}

func testIsValidUUID(t *testing.T) {
	assert := assert.New(t)

	assert.True(isValidUUID("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), "Valid UUID is returned as true")
	assert.False(isValidUUID("bad-id"), "Invalid UUID is returned as false")
}
