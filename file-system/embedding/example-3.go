package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	files, _ := fs.ReadDir(staticFiles, "static")
	for _, file := range files {
		fmt.Println("Found file:", file.Name())
	}
}
