# movie-gallery
GoLang REST API for querying movie information

# Docker commands
* `sudo docker build --tag docker-movie-gallery .`
* `sudo docker run -p 8080:8080 docker-movie-gallery`

# GraphQL mutations
```
mutation createMovie ($input: MovieInput!) {
  createMovie(input: $input) {
    name
    genre
    releaseYear
    director
    composer
    actors {
      name
    }
  }
}

mutation updateMovie ($movieId: String!, $input: MovieInput!) {
  updateMovie(movieId: $movieId, input: $input) {
    movieId
    name
    genre
    releaseYear
    director
    composer
    actors {
      name
    }
  }
}

mutation deleteMovie ($movieId: String!) {
  deleteMovie(movieId: $movieId)
}
```

# GraphQL query
```
query findAllMovies {
  findAllMovies {
    movieId
    name
    genre
    releaseYear
    director
    composer
    actors {
      name
    }
  }
}

query findMoviesByMovieName{
  findMoviesByField(name: "Tenet") {
    movieId
    name
    genre
    releaseYear
    director
    composer
    actors {
      name
    }
  }
}

query findMoviesByActorName{
  findMoviesByField(actors: ["Michael Caine"]) {
    movieId
    name
    genre
    releaseYear
    director
    composer
    actors {
      name
    }
  }
}

query findMoviesByComposite{
  findMoviesByField(composer: "Hans Zimmer", actors: ["Cillian Murphy"]) {
    movieId
    name
    genre
    releaseYear
    director
    composer
    actors {
      name
    }
  }
}
```

# GraphQL variables
```
{
  "input": {
    "name": "Tenet",
    "genre": "Action",
    "releaseYear": "2020",
    "director": "Christopher Nolan",
    "composer": "Ludwig Goransson",
    "actors": [
      {
        "name": "John David Washington"
      },
      {
        "name": "Robert Pattinson"
      },
      {
        "name": "Elizabeth Debicki"
      },
      {
        "name": "Dimple Kapadia"
      },
      {
        "name": "Michael Caine"
      },
      {
        "name": "Kenneth Branagh"
      }
    ]
  }
}
```

```
{
  "input": {
    "name": "Inception",
    "genre": "Action",
    "releaseYear": "2010",
    "director": "Christopher Nolan",
    "composer": "Hans Zimmer",
    "actors": [
      {
        "name": "Leonardo Dicaprio"
      },
      {
        "name": "Joseph Gordon-Levitt"
      },
      {
        "name": "Marion Cotillard"
      },
      {
        "name": "Eliot Page"
      },
      {
        "name": "Tom Hardy"
      },
      {
        "name": "Dileep Rao"
      },
      {
        "name": "Michael Caine"
      },
      {
        "name": "Cillian Murphy"
      }
    ]
  }
}
```

```
{
  "input": {
    "name": "Interstellar",
    "genre": "SciFi",
    "releaseYear": "2014",
    "director": "Christopher Nolan",
    "composer": "Hans Zimmer",
    "actors": [
      {
        "name": "Matthew McConaughey"
      },
      {
        "name": "Anne Hathaway"
      },
      {
        "name": "Jessica Chastain"
      },
      {
        "name": "Matt Damon"
      },
      {
        "name": "Michael Caine"
      }
    ]
  }
}
```

```
{
  "input": {
    "name": "The Dark Knight",
    "genre": "Superhero",
    "releaseYear": "2008",
    "director": "Christopher Nolan",
    "composer": "Hans Zimmer",
    "actors": [
      {
        "name": "Christian Bale"
      },
      {
        "name": "Michael Caine"
      },
      {
        "name": "Heath Ledger"
      },
      {
        "name": "Gary Oldman"
      },
      {
        "name": "Aaron Eckhart"
      },
      {
        "name": "Maggie Gyllenhaal"
      },
      {
        "name": "Morgan Freeman"
      }
    ]
  }
}
```
