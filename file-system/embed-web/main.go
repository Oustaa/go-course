package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var embeddedFiles embed.FS

func main() {
	// Create a filesystem from the embedded folder
	subFS, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	// Serve embedded files over HTTP
	fileServer := http.FileServer(http.FS(subFS))

	http.Handle("/", fileServer)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
