package gallery

import (
	"errors"
	models "github.com/harrlight00/movie-gallery/internal/gallery/models"
	"gorm.io/gorm/clause"
	"strings"
)

// Database function used for searching movies with any qualifiers provided as a
// MovieInfo object. Will only return up to 20 movies until pagination is implemented.
// TODO: add pagination
func getMovies(movieInfo *models.MovieInfo) ([]models.MovieInfo, error) {
	movies := make([]models.MovieInfo, 0)

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
		for _, actor := range movieInfo.Actors {
			fields = append(fields, "EXISTS(SELECT * FROM movie_actors _ma "+
				"INNER JOIN actors _a ON _a.id=_ma.actor_db_id "+
				"WHERE m.id = _ma.movie_db_id AND _a.name = ?)")
			values = append(values, actor)
		}
	}

	if len(fields) != 0 {
		query += " WHERE " + strings.Join(fields, " AND ")
	}

	query += " LIMIT 20;"

	rows, err := db.Raw(query, values...).Rows()
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
	movieInfo.MovieId = createUUID()
	movie := createMovieFromMovieInfo(*movieInfo)

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

// Database function used to grab a movie using a MovieId
func getMovie(movieId string) (models.MovieInfo, error) {
	var movie models.Movie
	if result := db.Select("Id", "MovieId", "Name", "Genre", "ReleaseYear", "Director", "Composer").
		Where(&models.Movie{MovieId: movieId}).First(&movie); result.Error != nil {
		return models.MovieInfo{}, result.Error
	}

	// Convert scanned movie into movieInfo
	movieInfo := createMovieInfoFromMovie(movie)

	// Fill in movieInfo.actors
	getMovieActors(&movieInfo)

	return movieInfo, nil
}

// Database function used for updating a current movie. The function will insert the movie,
// insert any new actors, and insert a movie_actor entry for each new movie_actor combo.
// Note: this function does not support deleting existing actors from a movie
func updateMovie(movieInfo *models.MovieInfo) error {
	// Convert MovieInfo to Movie
	movie := createMovieFromMovieInfo(*movieInfo)

	// Only update fields that have been specified
	fieldsToUpdate := make([]string, 0)
	if movieInfo.Name != "" {
		fieldsToUpdate = append(fieldsToUpdate, "Name")
	}
	if movieInfo.Genre != "" {
		fieldsToUpdate = append(fieldsToUpdate, "Genre")
	}
	if movieInfo.ReleaseYear != "" {
		fieldsToUpdate = append(fieldsToUpdate, "ReleaseYear")
	}
	if movieInfo.Director != "" {
		fieldsToUpdate = append(fieldsToUpdate, "Director")
	}
	if movieInfo.Composer != "" {
		fieldsToUpdate = append(fieldsToUpdate, "Composer")
	}

	if (len(fieldsToUpdate)) == 0 {
		return errors.New("No fields to update")
	}

	// Update movie (without actors)
	if result := db.Where("movie_id = ?", movie.MovieId).Updates(&movie); result.Error != nil {
		return result.Error
	}

	// Update actors
	for _, actorName := range movieInfo.Actors {
		var actor models.Actor
		// Check if Actor currently exists and create actor if not
		if result := db.Select("Id", "Name").
			FirstOrCreate(&actor, models.Actor{Name: actorName}); result.Error != nil {
			return result.Error
		}

		// Insert new MovieActor entry (if needed) with new/existing actor and new movie
		movieActor := models.MovieActor{MovieDbId: movie.Id, ActorDbId: actor.Id}
		if result := db.Clauses(clause.Insert{Modifier: "IGNORE"}).Select("MovieDbId", "ActorDbId").
			Create(&movieActor); result.Error != nil {
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

	rows, err := db.Raw(query, movieInfo.MovieId).Rows()
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
