// package main

// import (
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.StripPrefix("/statics/", http.FileServer(http.Dir("./public")))

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "./index.html")
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
