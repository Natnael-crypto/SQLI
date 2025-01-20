package controllers

import (
	"fmt"
	"net/http"
	"sqli/models"
	"sqli/views"
	"strconv"
)

func AdminController(w http.ResponseWriter, req *http.Request) {
	var (
		err      error
		products []models.Product
		isAdmin bool
	)

	isAdminCookie, err := req.Cookie("isAdmin")
	if err != nil {
		w.WriteHeader(403)
		views.ErrorRender(w, "403 Forbidden")
		return
	}

	isAdmin, _ = strconv.ParseBool(isAdminCookie.Value)
	if !isAdmin {
		w.WriteHeader(403)
		views.ErrorRender(w, "403 Forbidden")
		return
	}
	


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

	switch req.Method {
	case http.MethodGet:
		// Fetch all products for admin view
		products, err = models.GetAllProducts() // Assuming GetAllProducts fetches all products
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		adminVM := views.AdminVM{
			Username: username,
			Products: products,
		}

		// Render the admin page
		views.AdminRender(w, adminVM)

	case http.MethodPost:
		// Handle delete product
		if req.FormValue("action") == "delete" {
			productID := req.FormValue("product_id")
			fmt.Printf(productID)
			http.Redirect(w, req, "/admin", http.StatusFound)
		}

		// Handle update product
		if req.FormValue("action") == "update" {
			productID := req.FormValue("product_id")
			// Assuming you will redirect to a product update page
			http.Redirect(w, req, fmt.Sprintf("/update_product?id=%s", productID), http.StatusFound)
		}
	}
}
