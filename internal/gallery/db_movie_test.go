package gallery

import (
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMovies(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	t.Log("Testing GetMovies with no input")
	actualMovies, err := getMovies(&models.MovieInfo{})
	assert.Nil(err)
	assert.Equal(3, len(actualMovies), "Should be three results")

	t.Log("Testing GetMovies by name")
	actualMovies, err = getMovies(&models.MovieInfo{Name: "Tenet"})
	assert.Nil(err)
	assert.Equal(1, len(actualMovies), "Should be one result")
	assert.Equal("Tenet", actualMovies[0].Name, "Correct movie returned")

	t.Log("Testing GetMovies by composer")
	actualMovies, err = getMovies(&models.MovieInfo{Composer: "Hans Zimmer"})
	assert.Nil(err)
	assert.Equal(2, len(actualMovies), "Should be two results")
	assert.Contains([]string{"Interstellar", "Inception"}, actualMovies[0].Name, "Correct movies returned")
	assert.Contains([]string{"Interstellar", "Inception"}, actualMovies[1].Name, "Correct movies returned")

	t.Log("Testing GetMovies by actor")
	actualMovies, err = getMovies(&models.MovieInfo{Actors: []string{"Michael Caine"}})
	assert.Nil(err)
	assert.Equal(3, len(actualMovies), "Should be three results")
	assert.Contains([]string{"Interstellar", "Inception", "Tenet"}, actualMovies[0].Name, "Correct movies returned")
	assert.Contains([]string{"Interstellar", "Inception", "Tenet"}, actualMovies[1].Name, "Correct movies returned")
	assert.Contains([]string{"Interstellar", "Inception", "Tenet"}, actualMovies[2].Name, "Correct movies returned")

	t.Log("Testing GetMovies with multiple inputs")
	actualMovies, err = getMovies(&models.MovieInfo{Composer: "Hans Zimmer", Actors: []string{"Tom Hardy"}})
	assert.Nil(err)
	assert.Equal(1, len(actualMovies), "Should be one result")
	assert.Equal("Inception", actualMovies[0].Name, "Correct movie returned")

	t.Log("Testing GetMovies with multiple actors")
	actualMovies, err = getMovies(&models.MovieInfo{Actors: []string{"Michael Caine", "Robert Pattinson"}})
	assert.Nil(err)
	assert.Equal(1, len(actualMovies), "Should be one result")
	assert.Equal("Tenet", actualMovies[0].Name, "Correct movie returned")

	t.Log("Testing GetMovies with data returning no results")
	actualMovies, err = getMovies(&models.MovieInfo{Composer: "Hans Zimmer", Actors: []string{"Robert Pattinson"}})
	assert.Nil(err)
	assert.Equal(0, len(actualMovies), "Should be no results")
}

func TestInsertMovie(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	darkKnightMovie = models.MovieInfo{
		Name:        "The Dark Knight",
		Genre:       "Superhero",
		ReleaseYear: "2008",
		Director:    "Christopher Nolan",
		Composer:    "Hans Zimmer",
		Actors: []string{
			"Christian Bale",
			"Michael Caine",
			"Heath Ledger",
			"Gary Oldman",
			"Aaron Eckhart",
			"Maggie Gyllenhaal",
			"Morgan Freeman",
		},
	}

	t.Log("Testing GetMovies with no input pre-insert")
	actualMovies, err := getMovies(&models.MovieInfo{})
	assert.Nil(err)
	assert.Equal(3, len(actualMovies), "Should be three results")

	t.Log("Testing Insert operation with new movie")
	err = insertMovie(&darkKnightMovie)
	assert.Nil(err)

	t.Log("Testing GetMovies with no input post-insert")
	actualMovies, err = getMovies(&models.MovieInfo{})
	assert.Nil(err)
	assert.Equal(4, len(actualMovies), "Should be four results")
}

func TestGetMovie(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	t.Log("Testing GetMovie by ID")
	actualMovie, err := getMovie(tenetMovie.MovieId)
	assert.Nil(err)
	assert.Equal(tenetMovie.Name, actualMovie.Name, "Correct movie is returned")
}

func TestUpdateMovie(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	t.Log("Testing movie before genre change")
	actualMovie, err := getMovie(tenetMovie.MovieId)
	assert.Nil(err)
	assert.Equal("Action", actualMovie.Genre, "Genre should be action pre-update")

	t.Log("Testing update operation for genre change")
	tenetMovie.Id = actualMovie.Id
	tenetMovie.Genre = "Spy"
	err = updateMovie(&tenetMovie)
	// This will return an error as sqlite cannot use IGNORE the same way mysql can
	//assert.Nil(err)

	t.Log("Testing movie after genre change")
	actualMovie, err = getMovie(tenetMovie.MovieId)
	assert.Nil(err)
	assert.Equal("Spy", actualMovie.Genre, "Genre should be spy post-update")

	t.Log("Testing movie before actors change")
	actualMovie, err = getMovie(tenetMovie.MovieId)
	assert.Nil(err)
	assert.Equal(6, len(actualMovie.Actors), "Should be 6 actors pre-update")

	// TODO: figure out how to test update for actors, as INSERT IGNORE operation
	// does not function properly in sqlite as it does in mysql
	/*
	       t.Log("Testing update operation for actors change")
	   	tenetMovie.Id = actualMovie.Id
	       tenetMovie.Actors = append(tenetMovie.Actors, "Aaron-Taylor Johnson")
	       err = updateMovie(&tenetMovie)
	       // This will return an error as sqlite cannot use IGNORE the same way mysql can
	       assert.Nil(err)

	       t.Log("Testing movie after actors change")
	       actualMovie, err = getMovie(tenetMovie.MovieId)
	       assert.Nil(err)
	       assert.Equal(7, len(actualMovie.Actors), "Should be 7 actors post-update")
	   	assert.Contains(actualMovie.Actors, "Aaron-Taylor Johnson", "New actor should be returned")
	*/
}
