package initializers

import (
	"html/template"
	"io/fs"
	"log"
)

var Template *template.Template

func ParseTemplates(fileSystem fs.FS) {
	var err error
	Template, err = template.ParseFS(fileSystem, "assets/*.html")
	if err != nil {
		log.Printf("error occured while parsing templates, %v\n", err)
	}

}
