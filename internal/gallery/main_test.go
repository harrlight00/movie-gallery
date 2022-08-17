package gallery

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/movies", GetMovies)
	r.POST("/movies", CreateMovie)
	r.GET("/movies/:id", GetMovie)
	r.POST("/movies/:id", UpdateMovie)
	return r
}
