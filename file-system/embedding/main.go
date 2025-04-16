package main

import (
	"embed"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

var (
	//go:embed templates
	f embed.FS
)

func masdwin() {
	p := Person{"John", 27}
	tmpl, err := template.ParseFS(f, "templates/template.txt")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
