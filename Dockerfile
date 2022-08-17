# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./cmd/movie-gallery/*.go ./cmd/movie-gallery/
COPY ./dev_config.json ./
COPY ./internal/auth/*.go ./internal/auth/
COPY ./internal/config/*.go ./internal/config/
COPY ./internal/gallery/*.go ./internal/gallery/
COPY ./internal/gallery/models/*.go ./internal/gallery/models/
COPY ./internal/middleware/*.go ./internal/middleware/

RUN go build -o /docker-movie-gallery ./cmd/movie-gallery/main.go

EXPOSE 8080

CMD [ "/docker-movie-gallery" ]
