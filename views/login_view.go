package views

import (
	"io"
	"log"
	"sqli/initializers"
)

func LoginRender(file io.Writer) {
	err := initializers.Template.ExecuteTemplate(file, "login.gohtml", nil)
	if err != nil {
		log.Printf("error occured while executing template, %v\n", err)
	}

}
