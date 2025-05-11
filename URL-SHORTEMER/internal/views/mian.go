package views

import (
	"embed"
	"html/template"
	"io"
	"log"
	"net/http"
)

//go:embed templates/*.html
var templates embed.FS

type User struct {
	ID int
}

func MainView(w http.ResponseWriter, r *http.Request) {
	file, err := templates.Open("templates/main.html")
	if err != nil {
		log.Printf("error opening main.html file: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("error reading main.html content: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t, err := template.New("mainTemplate").Parse(string(content))
	if err != nil {
		log.Printf("error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, User{ID: 1})
	if err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
