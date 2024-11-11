package views

import (
	"io"
	"log"
	"text/template"
)

func LoginRender(file io.Writer) {
	templ, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		log.Fatalf("error while reading template, %v\n", err)
	}

	templ.ExecuteTemplate(file, "login.gohtml", nil)

}
