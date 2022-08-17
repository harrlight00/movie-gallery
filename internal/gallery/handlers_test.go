package gallery

import (
	"encoding/json"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenerateTokenHandler(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var tokenResponse TokenResponse

	t.Log("Test GenerateToken")
	responseData, responseCode := sendHttpRequest(t, r, nil, "POST", "/token")
	err := json.Unmarshal(responseData, &tokenResponse)
	assert.Nil(err)

	assert.Equal(http.StatusOK, responseCode, "200 response")
	assert.True(len(tokenResponse.Token) > 0, "Non-empty token created")
}

func TestGetMoviesHandler(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var actualMovies []models.MovieInfo

	t.Log("Test GetMovies by name")
	responseData, responseCode := sendHttpRequest(t, r, models.MovieInfo{Name: "Tenet"}, "GET", "/api/movies")
	err := json.Unmarshal(responseData, &actualMovies)
	assert.Nil(err)

	assert.Equal(http.StatusOK, responseCode, "200 response")
	assert.Equal(1, len(actualMovies), "1 result returned")
	assert.Equal("Tenet", actualMovies[0].Name, "Correct Result returned")
}

func TestCreateMovieHandler(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var actualMovie models.MovieInfo

	// darkKnightMovie is defined in main_test.go
	t.Log("Test InsertMovie success")
	// Marshal struct into JSON request data
	responseData, responseCode := sendHttpRequest(t, r, darkKnightMovie, "POST", "/api/movies")
	err := json.Unmarshal(responseData, &actualMovie)
	assert.Nil(err)

	assert.Equal(http.StatusOK, responseCode, "200 response")
	assert.Equal("The Dark Knight", actualMovie.Name, "Correct Result returned")
	assert.True(isValidUUID(actualMovie.MovieId))

	t.Log("Test InsertMovie fail with MovieId defined")
	darkKnightMovie.MovieId = "test_movie_id"
	// Marshal struct into JSON request data
	responseData, responseCode = sendHttpRequest(t, r, darkKnightMovie, "POST", "/api/movies")
	darkKnightMovie.MovieId = ""

	assert.Equal(http.StatusBadRequest, responseCode, "400 response")
	assert.Equal(`{"error":"Cannot create with MovieId field defined"}`,
		string(responseData), "Error matches expected")

	t.Log("Test InsertMovie fail with missing field")
	darkKnightMovie.Name = ""
	// Marshal struct into JSON request data
	responseData, responseCode = sendHttpRequest(t, r, darkKnightMovie, "POST", "/api/movies")
	darkKnightMovie.Name = "The Dark Knight"

	assert.Equal(http.StatusBadRequest, responseCode, "400 response")
	assert.Equal(`{"error":"Cannot create without Name field defined"}`,
		string(responseData), "Error matches expected")

	t.Log("Test InsertMovie fail with missing actors")
	darkKnightMovie.Actors = make([]string, 0)
	// Marshal struct into JSON request data
	responseData, responseCode = sendHttpRequest(t, r, darkKnightMovie, "POST", "/api/movies")

	assert.Equal(http.StatusBadRequest, responseCode, "400 response")
	assert.Equal(`{"error":"Cannot create without Actors field defined"}`,
		string(responseData), "Error matches expected")
}

func TestGetMovieHandler(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var actualMovie models.MovieInfo

	// tenetMovie is defined in main_test.go
	t.Log("Test GetMovie success")
	// Marshal struct into JSON request data
	responseData, responseCode := sendHttpRequest(t, r, nil, "GET", "/api/movies/"+tenetMovie.MovieId)
	err := json.Unmarshal(responseData, &actualMovie)
	assert.Nil(err)

	assert.Equal(http.StatusOK, responseCode, "200 response")
	assert.Equal("Tenet", actualMovie.Name, "Correct Result returned")

	t.Log("Test GetMovie fail with invalid MovieId")
	// Marshal struct into JSON request data
	responseData, responseCode = sendHttpRequest(t, r, nil, "GET", "/api/movies/bad-id")

	assert.Equal(http.StatusBadRequest, responseCode, "400 response")
	assert.Equal(`{"error":"Please input a valid MovieID"}`,
		string(responseData), "Error matches expected")
}

func TestUpdateMovieHandler(t *testing.T) {
	assert := assert.New(t)
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var actualMovie models.MovieInfo

	// tenetMovie is defined in main_test.go
	t.Log("Test UpdateMovie success")
	// Marshal struct into JSON request data
	tenetMovie.Genre = "Spy"
	responseData, responseCode := sendHttpRequest(t, r, tenetMovie, "POST", "/api/movies/"+tenetMovie.MovieId)
	err := json.Unmarshal(responseData, &actualMovie)
	assert.Nil(err)

	assert.Equal(http.StatusOK, responseCode, "200 response")
	assert.Equal("Tenet", actualMovie.Name, "Correct Result returned")
	assert.Equal("Spy", actualMovie.Genre, "Correct Genre returned")

	t.Log("Test UpdateMovie fail with invalid MovieId")
	// Marshal struct into JSON request data
	responseData, responseCode = sendHttpRequest(t, r, tenetMovie, "POST", "/api/movies/bad-id")

	assert.Equal(http.StatusBadRequest, responseCode, "400 response")
	assert.Equal(`{"error":"Please input a valid MovieID"}`,
		string(responseData), "Error matches expected")

	t.Log("Test UpdateMovie with non-matching MovieId")
	tenetMovie.MovieId = "bad-id"
	// Marshal struct into JSON request data
	responseData, responseCode = sendHttpRequest(t, r, tenetMovie, "POST", "/api/movies/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	tenetMovie.MovieId = ""

	assert.Equal(http.StatusBadRequest, responseCode, "400 response")
	assert.Equal(`{"error":"Please input matching MovieIDs in the URI and body"}`,
		string(responseData), "Error matches expected")

	t.Log("Test UpdateMovie with MovieId that doesn't map to a movie")
	// Marshal struct into JSON request data
	responseData, responseCode = sendHttpRequest(t, r, tenetMovie, "POST", "/api/movies/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")

	assert.Equal(http.StatusBadRequest, responseCode, "400 response")
	assert.Equal(`{"error":"MovieID does not map to an existing movie"}`,
		string(responseData), "Error matches expected")
}
