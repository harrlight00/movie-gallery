package gallery

import (
	"github.com/gin-gonic/gin"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"net/http"
)

// This function can be used for querying movies based on any field. This API will accept a
// request in the form of the Movie struct defined in the models section with any fields the
// user wishes to query for. As all fields are optional, any omitted fields will not be queried
// Results are limited to 20.
func GetMovies(c *gin.Context) {
	var movie models.MovieInfo

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	movies, err := getMovies(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// This function can be used for updating movies. This API will accept a request in the form
// of the Movie struct defined in the models folder with any fields the user wishes to update
func CreateMovie(c *gin.Context) {
	var movieInfo models.MovieInfo

	if err := c.ShouldBindJSON(&movieInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if movieInfo.MovieId != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create with ID field defined",
		})
		return
	}
	if movieInfo.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create without Name field defined",
		})
		return
	}
	if movieInfo.Genre == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create without Genre field defined",
		})
		return
	}
	if movieInfo.ReleaseYear == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create without ReleaseYear field defined",
		})
		return
	}
	if movieInfo.Director == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create without Director field defined",
		})
		return
	}
	if movieInfo.Composer == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create without Composer field defined",
		})
		return
	}
	if movieInfo.Actors == nil || len(movieInfo.Actors) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create without actors field defined",
		})
		return
	}

	err := insertMovie(&movieInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, movieInfo)
}

func GetMovie(c *gin.Context) {
	var param struct {
		MovieId string `uri:"id"`
	}

	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !isValidUUID(param.MovieId) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please input a valid MovieID",
		})
		return
	}

	movieInfo, err := getMovie(param.MovieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, movieInfo)
}

// This function can be used for updating movies. This API will accept a request in the form
// of the Movie struct defined in the models folder with any fields the user wishes to update
func UpdateMovie(c *gin.Context) {
	var movieInfo models.MovieInfo
	var param struct {
		MovieId string `uri:"id"`
	}

	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&movieInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !isValidUUID(param.MovieId) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please input a valid MovieID",
		})
		return
	}

	if movieInfo.MovieId != "" && movieInfo.MovieId != param.MovieId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please input matching MovieIDs in the URI and body",
		})
		return
	}

	dbMovie, validationErr := getMovie(param.MovieId)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "MovieID does not map to an existing movie",
		})
		return
	}

	// Set movieId and Id on the movieInfo object
	movieInfo.Id = dbMovie.Id
	movieInfo.MovieId = param.MovieId

	err := updateMovie(&movieInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, movieInfo)
}
