package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MessageData struct {
	Names []string `json:"fucking_names"`
}

func getDate() []string {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	messageJson := MessageData{}

	err = json.Unmarshal(data, &messageJson)
	if err != nil {

		log.Fatalln(err)
	}

	return messageJson.Names
}

func main() {
	data := getDate()

	fmt.Println(data)
}
