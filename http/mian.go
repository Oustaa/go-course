package main

// import (
// 	"log"
// 	"net/http"
// )

// type MyHandler struct {
// }

// func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	_, err := w.Write([]byte("HI"))
// 	if err != nil {
// 		log.Printf("an error occurred: %v\n", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// }

// // func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// // 	w.Write([]byte("Hello, World!"))
// // }

// func main() {
// 	log.Fatal(http.ListenAndServe(":8080", MyHandler{}))

// }
