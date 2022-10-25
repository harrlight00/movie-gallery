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

// Global token for use in handler testing
var test_token string

// Struct used for de-serializing token information
type TokenResponse struct {
	Token string
}

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

	gormDb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	assert.Nil(err)

	// Using global defined in main.go
	DB = gormDb

	// Create tables
	err = DB.AutoMigrate(&models.Movie{}, &models.MovieActor{}, &models.Actor{})
	assert.Nil(err)

	// Insert test records
	err = insertMovie(&tenetMovie)
	assert.Nil(err)
	err = insertMovie(&inceptionMovie)
	assert.Nil(err)
	err = insertMovie(&interstellarMovie)
	assert.Nil(err)

	// Generate router
	r = SetUpRouter()

	// Generate token
	var tokenResponse TokenResponse
	responseData, responseCode := sendHttpRequest(t, r, nil, "POST", "/token")
	err = json.Unmarshal(responseData, &tokenResponse)
	assert.Nil(err)

	assert.Equal(http.StatusOK, responseCode, "200 response")
	test_token = tokenResponse.Token
	t.Log("Generated token: " + test_token)
	assert.True(len(test_token) > 0, "Non-empty token created")

	return func(t *testing.T) {
		t.Log("Teardown test case")
		DB.Exec("DELETE FROM movie_actors")
		DB.Exec("DELETE FROM actors")
		DB.Exec("DELETE FROM movies")
	}
}

// Helper method used for testing HTTP requests
func sendHttpRequest(t *testing.T, r *gin.Engine, data interface{}, httpMethod string, url string) ([]byte, int) {
	assert := assert.New(t)

	requestJSON, err := json.Marshal(data)
	assert.Nil(err)

	// Test request
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestJSON))
	assert.Nil(err)

	// Add headers
	req.Header.Set("Authorization", test_token)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Unmarshal JSON into struct
	responseData, err := ioutil.ReadAll(w.Body)
	assert.Nil(err)
	return responseData, w.Code
}
