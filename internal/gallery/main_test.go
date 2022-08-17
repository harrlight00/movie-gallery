package gallery

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tenetMovie models.MovieInfo
var inceptionMovie models.MovieInfo
var interstellarMovie models.MovieInfo
var darkKnightMovie models.MovieInfo

// Helper method used for setting and clearing context before and after tests
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("Setup test case with DB contents")

	// Instantiate test movies
	tenetMovie = models.MovieInfo{
		Name:        "Tenet",
		Genre:       "Action",
		ReleaseYear: "2020",
		Director:    "Christopher Nolan",
		Composer:    "Ludwig Goransson",
		Actors: []string{
			"John David Washington",
			"Robert Pattinson",
			"Elizabeth Debicki",
			"Dimple Kapadia",
			"Michael Caine",
			"Kenneth Branagh",
		},
	}

	inceptionMovie = models.MovieInfo{
		Name:        "Inception",
		Genre:       "Action",
		ReleaseYear: "2010",
		Director:    "Christopher Nolan",
		Composer:    "Hans Zimmer",
		Actors: []string{
			"Leonardo Dicaprio",
			"Joseph Gordon-Levitt",
			"Marion Cotillard",
			"Eliot Page",
			"Tom Hardy",
			"Dileep Rao",
			"Michael Caine",
			"Cillian Murphy",
		},
	}

	interstellarMovie = models.MovieInfo{
		Name:        "Interstellar",
		Genre:       "SciFi",
		ReleaseYear: "2014",
		Director:    "Christopher Nolan",
		Composer:    "Hans Zimmer",
		Actors: []string{
			"Matthew McConaughey",
			"Anne Hathaway",
			"Jessica Chastain",
			"Matt Damon",
			"Michael Caine",
		},
	}

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

	assert := assert.New(t)

	gormDb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.Nil(err)

	// Using global defined in main.go
	db = gormDb

	// Create tables
	err = db.AutoMigrate(&models.Movie{}, &models.MovieActor{}, &models.Actor{})
	assert.Nil(err)

	// Insert test records
	err = insertMovie(&tenetMovie)
	assert.Nil(err)
	err = insertMovie(&inceptionMovie)
	assert.Nil(err)
	err = insertMovie(&interstellarMovie)
	assert.Nil(err)

	return func(t *testing.T) {
		t.Log("Teardown test case")
		db.Exec("DELETE FROM movie_actors")
		db.Exec("DELETE FROM actors")
		db.Exec("DELETE FROM movies")
	}
}

// Helper method used for setting up router for tests
func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/movies", GetMovies)
	r.POST("/movies", CreateMovie)
	r.GET("/movies/:id", GetMovie)
	r.POST("/movies/:id", UpdateMovie)
	return r
}

// Helper method used for testing HTTP requests
func sendHttpRequest(t *testing.T, r *gin.Engine, data interface{}, httpMethod string, url string) ([]byte, int) {
	assert := assert.New(t)

	requestJSON, err := json.Marshal(data)
	assert.Nil(err)

	// Test request
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestJSON))
	assert.Nil(err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Unmarshal JSON into struct
	responseData, err := ioutil.ReadAll(w.Body)
	assert.Nil(err)
	return responseData, w.Code
}
