package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	counter int
	content string
	heading string
}

func (pwc *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pwc.counter++

	pageContent := fmt.Sprintf(`
		<h1>%s</h1>
		%s
		<p>total page views are: %d</p>
	`, pwc.heading, pwc.content, pwc.counter)

	_, err := w.Write([]byte(pageContent))
	if err != nil {
		log.Fatalf("There was an error serving page content, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {

	http.Handle("/", &PageWithCounter{counter: 0, heading: "Page with counter", content: "i am oussama tailba"})
	http.Handle("/chapter1", &PageWithCounter{counter: 0, heading: "Chapter 1", content: "this is my first chapter"})
	http.Handle("/chapter2", &PageWithCounter{counter: 0, heading: "Chapter 2", content: "this is my second chapter"})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
