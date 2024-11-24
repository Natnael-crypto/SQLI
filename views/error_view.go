package views

import (
	"io"
	"log"
	"sqli/initializers"
)

func ErrorRender(file io.Writer, errorMsg string) {
	err := initializers.Template.ExecuteTemplate(file, "error.html", errorMsg)
	if err != nil {
		log.Printf("error occured while executing template, %v\n", err)

	}
}