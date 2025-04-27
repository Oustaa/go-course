package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MessageData struct {
	Name string
}

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	messageData := MessageData{}

	err := json.NewDecoder(r.Body).Decode(&messageData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Println(messageData)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
