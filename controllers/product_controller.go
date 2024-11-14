package controllers

import (
	"fmt"
	"net/http"
	"sqli/models"
	"sqli/views"
)

func ProductsController(w http.ResponseWriter, req *http.Request) {
	var (
		err        error
		category   string
		products   []models.Product
		productVMs []models.ProductVM
	)
	
	req.ParseForm()
	action := req.FormValue("action")

	values := req.URL.Query()
	if len(values) > 0 {
		category = values["category"][0]
	}

	if action == Vuln {
		products, err = models.VulnGetProductsByCategory(category)
	} else {
		products, err = models.SecureGetProductsByCategory(category)
	}
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		for _, product := range products {
			productVM := product.GenerateViewModel()
			productVMs = append(productVMs, productVM)
		}

		views.ProductsRender(w, productVMs)
	}
}
