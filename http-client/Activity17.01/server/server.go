package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := struct {
		Names []string `json:"fucking_names"`
	}{Names: []string{"Electric", "Electric", "Electric", "Boogaloo", "Booga-loo", "Boogaloo", "Boogaloo"}}

	messageStr, err := json.Marshal(message)
	if err != nil {
		w.Write([]byte("An error parsin"))
		return
	}
	fmt.Println(string(messageStr))

	w.Write((messageStr))
}

func main() {
	log.Fatalln(http.ListenAndServe(":8080", Server{}))
}
