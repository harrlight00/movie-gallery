package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/harrlight00/movie-gallery/internal/graph/generated"
	"github.com/harrlight00/movie-gallery/internal/graph/model"
	"github.com/harrlight00/movie-gallery/internal/store"
)

// CreateMovie is the resolver for the createMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (*model.Movie, error) {
	movie := model.Movie{
		MovieID:     (uuid.New()).String(),
		Name:        input.Name,
		Genre:       input.Genre,
		ReleaseYear: input.ReleaseYear,
		Director:    input.Director,
		Composer:    input.Composer,
		Actors:      mapActorsFromInput(input.Actors),
	}

	if err := store.DB.Create(&movie).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

// UpdateMovie is the resolver for the updateMovie field.
func (r *mutationResolver) UpdateMovie(ctx context.Context, movieID string, input model.MovieInput) (*model.Movie, error) {
	movie := model.Movie{
		MovieID:     movieID,
		Name:        input.Name,
		Genre:       input.Genre,
		ReleaseYear: input.ReleaseYear,
		Director:    input.Director,
		Composer:    input.Composer,
		Actors:      mapActorsFromInput(input.Actors),
	}

	err := store.DB.Model(&model.Movie{}).Where("movie_id = ?", movieID).Updates(&movie).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

// DeleteMovie is the resolver for the deleteMovie field.
func (r *mutationResolver) DeleteMovie(ctx context.Context, movieID string) (bool, error) {
	var movie *model.Movie

	// Find movie
	store.DB.Where("movie_id = ?", movieID).First(&movie)

	// Delete associated actors and movie
	store.DB.Select("Actors").Delete(&movie)

	return true, nil
}

// FindAllMovies is the resolver for the findAllMovies field.
func (r *queryResolver) FindAllMovies(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie
	store.DB.Preload("Actors").Find(&movies)

	return movies, nil
}

// FindAllActors is the resolver for the findAllActors field.
func (r *queryResolver) FindAllActors(ctx context.Context) ([]*model.Actor, error) {
	var actors []*model.Actor
	store.DB.Find(&actors)

	return actors, nil
}

// OneMovie is the resolver for the oneMovie field.
func (r *queryResolver) OneMovie(ctx context.Context, movieID string) (*model.Movie, error) {
	var movie *model.Movie
	store.DB.Preload("Actors").Where("movie_id = ?", movieID).First(&movie)

	return movie, nil
}

// FindMoviesByField is the resolver for the findMoviesByField field.
func (r *queryResolver) FindMoviesByField(ctx context.Context, movieID *string, name *string, actors []*string, genre *string, releaseYear *string, director *string, composer *string) ([]*model.Movie, error) {
	var movies []*model.Movie

	// Write query based on which fields are included
	query := `SELECT DISTINCT m.movie_id, m.name, m.genre, m.release_year, m.director,
        m.composer FROM movies m LEFT JOIN actors a ON a.movie_db_id=m.id`

	fields := make([]string, 0)
	values := make([]interface{}, 0)

	if movieID != nil {
		fields = append(fields, "m.movie_id = ?")
		values = append(values, *movieID)
	}
	if name != nil {
		fields = append(fields, "m.name = ?")
		values = append(values, *name)
	}
	if genre != nil {
		fields = append(fields, "m.genre = ?")
		values = append(values, *genre)
	}
	if releaseYear != nil {
		fields = append(fields, "m.release_year = ?")
		values = append(values, *releaseYear)
	}
	if director != nil {
		fields = append(fields, "m.director = ?")
		values = append(values, *director)
	}
	if composer != nil {
		fields = append(fields, "m.composer = ?")
		values = append(values, *composer)
	}
	if actors != nil && len(actors) != 0 {
		for _, actor := range actors {
			fields = append(fields, `EXISTS(SELECT * FROM actors _a
                WHERE _a.movie_db_id = m.id AND _a.name = ?)`)
			values = append(values, actor)
		}
	}

	if len(fields) != 0 {
		query += " WHERE " + strings.Join(fields, " AND ")
	}

	query += " LIMIT 20;"

	rows, err := store.DB.Raw(query, values...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var movieResult *model.Movie
		store.DB.ScanRows(rows, &movieResult)

		var movie *model.Movie
		store.DB.Preload("Actors").Where("movie_id = ?", movieResult.MovieID).First(&movie)

		movies = append(movies, movie)
	}

	return movies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func mapActorsFromInput(actorsInput []*model.ActorInput) []*model.Actor {
	var actors []*model.Actor
	for _, actorInput := range actorsInput {
		actors = append(actors, &model.Actor{
			Name: actorInput.Name,
		})
	}
	return actors
}
