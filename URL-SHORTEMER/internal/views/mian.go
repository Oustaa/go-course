package views

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates/main.html
var mainTemplate string

type User struct {
	ID int
}

func MainView(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("mainTemplate").Parse(mainTemplate)

	err := t.Execute(w, User{1})
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
