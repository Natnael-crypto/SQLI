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

	userCookie, err := req.Cookie("User")

	if err != nil {
		// If the cookie is not found, handle the error
		if err == http.ErrNoCookie {
			http.Error(w, "Username cookie not found", http.StatusUnauthorized)
			return
		}
		// Handle any other potential errors
		http.Error(w, fmt.Sprintf("Error retrieving cookie: %v", err), http.StatusInternalServerError)
		return
	}

	username := userCookie.Value
	userVM := views.UserVM{
		Username: username,
		Products: productVMs,
	}

	req.ParseForm()
	action := req.FormValue("action")
	if action != "" {
		values := req.URL.Query()
		if len(values["category"]) > 0 {
			category = values["category"][0]
		}

		if action == Vuln {
			fmt.Printf("passed")
			// products =
			err = http.ErrNoCookie
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

			userVM := views.UserVM{
				Username: username,
				Products: productVMs,
			}

			views.ProductsRender(w, userVM)
		}
	} else {
		views.ProductsRender(w, userVM)
	}
}
