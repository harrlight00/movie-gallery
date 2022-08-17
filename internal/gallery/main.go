package gallery

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	config "github.com/harrlight00/movie-gallery/internal/config"
)

var db *gorm.DB

func StartServer() {
	// Connect to MySQL DB
	fmt.Println(dsn())
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

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/movies", GetMovies)
	r.POST("/movies", CreateMovie)
	r.GET("/movies/:id", GetMovie)
	r.POST("/movies/:id", UpdateMovie)

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
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
