package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sqli/models"
	"sqli/views"
)

func ProductsController(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	log.Printf("values: %v\n", values)
	var category string
	if len(values) > 0 {
		category = values["category"][0]
	}

	products, err := models.VulnGetProductsByCategory(category)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		productVMs := []models.ProductVM{}
		for _, product := range products {
			productVM := product.GenerateViewModel()
			productVMs = append(productVMs, productVM)
		}

		views.ProductsRender(w, productVMs)
	}
}
