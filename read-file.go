package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed files/*.go
var files embed.FS

func main() {
	file, err := files.Open("files/file.go")
	if err != nil {
		fmt.Printf("error opening the file, %v\n", err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error reading from the file, %v\n", err)
	}
	// var content string
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Bytes())
	// }

	fmt.Println(string(content))
}
