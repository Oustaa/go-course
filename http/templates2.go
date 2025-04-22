package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var html = `
	<html>
	<body>
		{{ if .ID }} 
			<h1>Welcome back, {{ .Username }}</h1>
			<h2>email: {{ .Email }}</h2>
		{{ else }}
			Please give us an id
		{{ end }}
	</body>
	</html>
`

type User struct {
	ID       int
	Username string
	Email    string
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		_, ok := params["id"]
		if ok {
			next.ServeHTTP(w, r)
			return
		}

		w.Write([]byte("You are not authorized"))
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	user := User{}

	id, ok := params["id"]
	if ok {
		user.ID, _ = strconv.Atoi(strings.Join(id, ","))
	}

	username, ok := params["username"]
	if ok {
		user.Username = strings.Join(username, "")
	}

	email, ok := params["email"]
	if ok {
		user.Email = strings.Join(email, "")
	}

	tmpl, _ := template.New("test").Parse(html)

	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", Auth(Hello))
	println("Server running at http://localhost:9000")

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err) // or: log.Fatal(err)
	}
}
