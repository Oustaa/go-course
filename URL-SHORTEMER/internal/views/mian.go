package views

import (
	_ "embed"
	"net/http"
)

//go:embed templates/main.html
var mainTemplate string

func MainView(w http.ResponseWriter, r *http.Request) {

}
