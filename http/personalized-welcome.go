package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func He2llo(w http.ResponseWriter, r *http.Request) {

	vl := r.URL.Query()

	name, ok := vl["name"]

	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name"))
		return
	}

	fmt.Println(strings.Join(name, "-"))
	fmt.Println(name)

	w.Write([]byte(fmt.Sprintf("Hello %s", strings.Join(name, ","))))
}
func mai2n() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
