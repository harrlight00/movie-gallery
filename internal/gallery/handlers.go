package gallery

import (
	"github.com/gin-gonic/gin"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"net/http"
)

// This function can be used for querying movies based on any field. This API will accept a
// request in the form of the Movie struct defined in the models section with any fields the
// user wishes to query for. As all fields are optional, any omitted fields will not be queried
// Results are limited to 10.
func GetMovies(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindQuery(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := getMovies(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// This function can be used for updating movies. This API will accept a request in the form
// of the Movie struct defined in the models folder with any fields the user wishes to update
func UpdateMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if movie.Id == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Field Id needs to be present in update request",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetMovie(c *gin.Context) {
	var param struct {
		Name string `uri:"id"`
	}

	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// This function can be used for updating movies. This API will accept a request in the form
// of the Movie struct defined in the models folder with any fields the user wishes to update
func UpdateMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if movie.Id == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot update without ID field",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
