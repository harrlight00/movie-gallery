package gallery

import (
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"strings"
)

// Database function used for searching movies with any qualifiers provided as a 
// MovieInfo object. Will only return up to 20 movies until pagination is implemented.
// Note: currently getMovies only accepts searching by one actor at a time
// TODO: add pagination
func getMovies(movieInfo *models.MovieInfo) ([]models.MovieInfo, error) {
	var movies []models.MovieInfo

	// Write query based on which fields are included
	query := `SELECT DISTINCT m.movie_id, m.name, m.genre, m.release_year, m.director, 
		m.composer FROM movie_actors ma INNER JOIN movies m ON m.id=ma.movie_db_id 
		INNER JOIN actors a ON a.id=ma.actor_db_id`

	fields := make([]string, 0)
	values := make([]interface{}, 0)

	if movieInfo.MovieId != "" {
		fields = append(fields, "m.movie_id = ?")
		values = append(values, movieInfo.MovieId)
	}
	if movieInfo.Name != "" {
		fields = append(fields, "m.name = ?")
		values = append(values, movieInfo.Name)
	}
	if movieInfo.Genre != "" {
		fields = append(fields, "m.genre = ?")
		values = append(values, movieInfo.Genre)
	}
	if movieInfo.ReleaseYear != "" {
		fields = append(fields, "m.release_year = ?")
		values = append(values, movieInfo.ReleaseYear)
	}
	if movieInfo.Director != "" {
		fields = append(fields, "m.director = ?")
		values = append(values, movieInfo.Director)
	}
	if movieInfo.Composer != "" {
		fields = append(fields, "m.composer = ?")
		values = append(values, movieInfo.Composer)
	}
	if movieInfo.Actors != nil && len(movieInfo.Actors) != 0 {
		fields = append(fields, "a.name = ?")
		values = append(values, movieInfo.Actors[0])
	}

	if (len(fields) != 0) {
		query += " WHERE " + strings.Join(fields, " AND ")
	}

	query += " LIMIT 20;"

	rows, err := db.Raw(query, values...).Rows();
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var movieResult models.Movie
		db.ScanRows(rows, &movieResult)

		// Convert scanned movie into movieInfo
		movieInfo := createMovieInfoFromMovie(movieResult)

		// Fill in movieInfo.actors
		getMovieActors(&movieInfo)

		movies = append(movies, movieInfo)
	}

	return movies, nil
}

// Database function used for inserting a new movie. The function will insert the movie, 
// insert any new actors, and insert a movie_actor entry for each new movie_actor combo.
func insertMovie(movieInfo *models.MovieInfo) error {
	// Convert MovieInfo to Movie
	movie := createMovieFromMovieInfo(*movieInfo)
	movie.MovieId = createUUID()

	// Create movie (without actors)
	if result := db.Select("MovieId", "Name", "Genre", "ReleaseYear", "Director", "Composer").
			Create(&movie); result.Error != nil {
		return result.Error
	}

	// Insert actors
	for _, actorName := range movieInfo.Actors {
		var actor models.Actor
		// Check if Actor currently exists and create actor if not
		if result := db.Select("Id", "Name").
				FirstOrCreate(&actor, models.Actor{Name: actorName}); result.Error != nil {
			return result.Error
		}

		// Insert new MovieActor entry with new/existing actor and new movie
		movieActor := models.MovieActor{MovieDbId: movie.Id, ActorDbId: actor.Id}
		if result := db.Select("MovieDbId", "ActorDbId").Create(&movieActor); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

// Database function used to fill in actor information for a MovieInfo object
func getMovieActors(movieInfo *models.MovieInfo) error {
	// Search MovieActors table
    query := `SELECT DISTINCT a.name FROM movie_actors ma INNER JOIN movies m ON m.id=ma.movie_db_id
        INNER JOIN actors a ON a.id=ma.actor_db_id WHERE m.movie_id = ?`

    rows, err := db.Raw(query, movieInfo.MovieId).Rows();
    if err != nil {
        return err
    }
    defer rows.Close()
    for rows.Next() {
        var actor models.Actor
        db.ScanRows(rows, &actor)

		// Update actors value with each actor
		movieInfo.Actors = append(movieInfo.Actors, actor.Name)
    }

	return nil
}
