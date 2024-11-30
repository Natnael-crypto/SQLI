package views

import (
	"io"
	"log"
	"sqli/initializers"
	"sqli/models"
)

type AdminVM struct {
	Username    string
	Products    []models.Product
}

func AdminRender(file io.Writer, adminVM AdminVM) {
	err := initializers.Template.ExecuteTemplate(file, "admin.html", adminVM)
	if err != nil {
		log.Printf("error occurred while executing template: %v\n", err)
	}
}
