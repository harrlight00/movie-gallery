package main

import (
	"io"
	"log"
	"net/http"
    gallery "github.com/harrlight00/movie-gallery/internal/gallery"
)

func main() {
    gallery.StartServer()
}
