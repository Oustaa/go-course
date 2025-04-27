package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type MessageData struct {
	Name string
}

func SendDataToServer(name string) {
	message := MessageData{Name: name}
	messageStr, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	r, err := http.Post("http://localhost:8080", "Application/json", bytes.NewBuffer(messageStr))

	r.Body.Close()
}

func main() {
	SendDataToServer("Oussama Tailba")
	SendDataToServer("Kaoutra Taki")
}
