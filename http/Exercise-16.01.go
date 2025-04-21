package main

// import (
// 	"log"
// 	"net/http"
// )

// type MyHandler struct{}

// func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	msg := "<h1>Hello World</h1>"
// 	_, err := w.Write([]byte(msg))
// 	if err != nil {
// 		log.Printf("an error occurred: %v\n", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// }

// func main() {
// 	http.HandleFunc("/my-name", func(w http.ResponseWriter, r *http.Request) {
// 		msg := "<h1>Oussama Tailba</h1><br/><h2>I am a full satck dev</h2>"
// 		_, err := w.Write([]byte(msg))
// 		if err != nil {
// 			log.Printf("an error occurred: %v\n", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 	})
// 	log.Fatal(http.ListenAndServe(":8080", MyHandler{}))
// }
