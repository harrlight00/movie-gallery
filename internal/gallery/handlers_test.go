package gallery

import (
	"bytes"
	"encoding/json"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMoviesHandler(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	assert.Equal(1, 1, "test")
	_ = models.MovieInfo{}

	r := SetUpRouter()

	var actualMovies []models.MovieInfo

	t.Log("Test GetMovies by name")
	// Marshal struct into JSON request data
	requestJSON, err := json.Marshal(models.MovieInfo{Name: "Tenet"})
	assert.Nil(err)

	// Test request
	req, err := http.NewRequest("GET", "/movies", bytes.NewBuffer(requestJSON))
	assert.Nil(err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Unmarshal JSON into struct
	responseData, err := ioutil.ReadAll(w.Body)
	assert.Nil(err)
	err = json.Unmarshal(responseData, &actualMovies)

	assert.Equal(http.StatusOK, w.Code, "200 response")
	assert.Equal(1, len(actualMovies), "1 result returned")
	assert.Equal("Tenet", actualMovies[0].Name, "Correct Result returned")
}
