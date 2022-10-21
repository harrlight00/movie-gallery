package gallery

import (
	"github.com/gin-gonic/gin"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	middleware "github.com/harrlight00/movie-gallery/internal/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Global used for accessing the DB
var db *gorm.DB

// Global used for accessing the router
var r *gin.Engine

// Method to start the HTTP server by creating the router and DB
func StartServer() {
	gormDb, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = gormDb

	if err := db.AutoMigrate(&models.Movie{}, &models.MovieActor{}, &models.Actor{}); err != nil {
		panic(err)
	}

	r = SetUpRouter()

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}

// Helper method used for setting up router
func SetUpRouter() *gin.Engine {
	r = gin.Default()
	r.SetTrustedProxies(nil)
	r.POST("/token", GenerateToken)
	api := r.Group("/api").Use(middleware.Auth())
	{
		api.GET("/movies", GetMovies)
		api.POST("/movies", CreateMovie)
		api.GET("/movies/:id", GetMovie)
		api.POST("/movies/:id", UpdateMovie)
	}
	return r
}
