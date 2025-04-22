package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tplStr = `
<html>
  <h1>Customer {{.ID}}</h1>
  {{if .ID }}
   <p>Details:</p>
   <ul>
   {{if .Name}}<li>Name: {{.Name}}</li>{{end}}
   {{if .Surname}}<li>Surname: {{.Surname}}</li>{{end}}
   {{if .Age}}<li>Age: {{.Age}}</li>{{end}}
   </ul>
  {{else}}
  <p>Data not available</p>
  {{end}}
</html>
`

type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

func Hel2lo(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()

	cust := Customer{}

	id, ok := vl["id"]
	if ok {
		cust.ID, _ = strconv.Atoi(strings.Join(id, ","))
	}
	name, ok := vl["name"]
	if ok {
		cust.Name = strings.Join(name, ",")
	}
	surname, ok := vl["surname"]
	if ok {
		cust.Surname = strings.Join(surname, ",")
	}
	age, ok := vl["age"]
	if ok {
		cust.Age, _ = strconv.Atoi(strings.Join(age, ""))
	}

	tmpl, _ := template.New("test").Parse(tplStr)

	tmpl.Execute(w, cust)
}

func mai22n() {
	http.HandleFunc("/", Hel2lo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
