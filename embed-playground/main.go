package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed files/*
var f embed.FS

func main() {
	// files, err := f.ReadDir(".")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	file, err := f.Open("files/index.html")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	println(string(fileContent))

	// var fileContent []byte
	// byteRead, err := file.Read(fileContent)

	// fmt.Printf("bytes read are: %d\n", byteRead)
	// fmt.Printf("i read:\n%s\n", fileContent)
	// fmt.Println(fileContent)

	// fmt.Println("Hello me")
}
