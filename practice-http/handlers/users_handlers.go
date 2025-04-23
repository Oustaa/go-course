package handlers

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/oustaa/practice-http/database"
)

//go:embed user.html
var userPage string

type UserHandler struct{}

func DisplayUserById(w http.ResponseWriter, r *http.Request) {
	var user database.User

	t, err := template.New("userPage").Parse(userPage)
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		log.Printf("Template parsing error: %v\n", err)
		return
	}

	params := r.URL.Query()
	pId, ok := params["id"]
	if !ok {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		log.Println("Missing user ID")
		return
	}

	id, err := strconv.ParseUint(strings.Join(pId, ","), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Invalid ID: %v\n", err)
		return
	}

	user, err = user.FindById(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		log.Printf("Error fetching user: %v\n", err)
		return
	}

	err = t.Execute(w, user)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Printf("Template execution error: %v\n", err)
	}
}
