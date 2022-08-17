package gallery

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	config "github.com/harrlight00/movie-gallery/internal/config"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	middleware "github.com/harrlight00/movie-gallery/internal/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Global used for accessing the DB
var db *gorm.DB

// Global used for accessing the router
var r *gin.Engine

// Method to start the HTTP server by creating the router and DB
func StartServer() {
	// Connect to MySQL DB
	mysqldb, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer mysqldb.Close()

	// Access SQL DB through Gorm
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqldb,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = gormDB

	if err := db.AutoMigrate(&models.Movie{}, &models.MovieActor{}, &models.Actor{}); err != nil {
		panic(err)
	}

	r = SetUpRouter()

	if err := r.Run("localhost:8080"); err != nil {
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

func dsn() string {
	dbConfig := config.GetConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.DB_USERNAME,
		dbConfig.DB_PASSWORD,
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_NAME,
	)
}
