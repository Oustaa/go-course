package routes

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/oustaa/url-shortner/internal/views"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	FileServer(r, "/static", http.Dir("./public"))

	r.Get("/", views.MainView)

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	// Detect development environment
	isDev := os.Getenv("ENV") == "development"

	fmt.Println(os.Getenv("ME"))

	r.Get(path+"/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDev {
			w.Header().Set("Cache-Control", "no-store")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}
		fs.ServeHTTP(w, r)
	}))
}
