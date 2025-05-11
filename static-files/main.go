// // package main

// // import (
// // 	"embed"
// // 	"fmt"
// // 	"io/fs"
// // 	"log"
// // 	"net/http"
// // )

// // //go:embed public/*
// // var files embed.FS

// // func main() {
// // 	subFS, err := fs.Sub(files, "public")
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	entries, err := fs.ReadDir(subFS, ".")
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	for _, e := range entries {
// // 		fmt.Println(e.Name())
// // 	}

// // 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(subFS))))

// //		fmt.Println("Server running at http://localhost:8080")
// //		log.Fatalln(http.ListenAndServe(":8000", nil))
// //	}
// package main

// import (
// 	"embed"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// //go:embed public/*
// var files embed.FS

// func main() {

// 	fs := http.FileServer(http.Dir("public"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	fmt.Println("Server running at http://localhost:8080")
// 	log.Fatalln(http.ListenAndServe(":8000", nil))
// }

package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatalln(http.ListenAndServe(":9000", nil))

}
