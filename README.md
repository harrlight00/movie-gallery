# movie-gallery
GoLang REST API for querying movie information

# Docker commands
* `sudo docker build --tag docker-movie-gallery .`
* `sudo docker run -p 8080:8080 docker-movie-gallery`

# Quickstart guide
* Please create a file at the root level titled dev_config.js to store a value for [JWT_KEY](https://github.com/harrlight00/movie-gallery/blob/main/internal/auth/jwtManager.go#L11) and [database info](https://github.com/harrlight00/movie-gallery/blob/main/internal/gallery/main.go#L66) if backing with a MySQL database.
* If you do not wish to use MySQL, please switch to a sqlite database [here](https://github.com/harrlight00/movie-gallery/blob/main/internal/gallery/main.go#L24) mirroring [test code](https://github.com/harrlight00/movie-gallery/blob/main/internal/gallery/main_test.go#L103)
