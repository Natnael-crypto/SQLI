package views

import (
	"io"
	"log"
	"sqli/initializers"
	"sqli/models"
)

func ProductsRender(file io.Writer, product []models.ProductVM) {
	err := initializers.Template.ExecuteTemplate(file, "products.html", product)
	if err != nil {
		log.Printf("error occured while executing template, %v\n", err)

	}
}
