package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getDataAndReturnResponse() string {
	r, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func ma23in() {
	data := getDataAndReturnResponse()

	file, err := os.OpenFile("index.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	file.Write([]byte(data))

	log.Println(data)
}
