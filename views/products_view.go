package views

import (
	"io"
	"log"
	"sqli/initializers"
	"sqli/models"
)

type UserVM struct {
	Username string
	Products []models.ProductVM
}

func ProductsRender(file io.Writer, userVM UserVM) {
	err := initializers.Template.ExecuteTemplate(file, "products.html", userVM)
	if err != nil {
		log.Printf("error occured while executing template, %v\n", err)

	}
}
