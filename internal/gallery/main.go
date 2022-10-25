package gallery

import (
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/harrlight00/movie-gallery/internal/graph"
	"github.com/harrlight00/movie-gallery/internal/graph/generated"
	"github.com/harrlight00/movie-gallery/internal/middleware"
	"github.com/harrlight00/movie-gallery/internal/store"
)

// Global used for accessing the DB
var DB *gorm.DB

// Global used for accessing the router
var r *gin.Engine

// Method to start the HTTP server by creating the router and DB
func StartServer() {
	store.ConnectDB()
	DB = store.DB

	r = SetUpRouter()

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Helper method used for setting up router
func SetUpRouter() *gin.Engine {
	r = gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/ping", Ping)
	r.POST("/token", GenerateToken)
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	api := r.Group("/api").Use(middleware.Auth())
	{
		api.GET("/movies", GetMovies)
		api.POST("/movies", CreateMovie)
		api.GET("/movies/:id", GetMovie)
		api.POST("/movies/:id", UpdateMovie)
	}
	return r
}
