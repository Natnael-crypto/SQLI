package main

import (
	"net/http"

	"sqli/controllers"
)

func Router() {
	http.HandleFunc("/login", controllers.LoginController)
	http.HandleFunc("/", controllers.LoginController)
	http.HandleFunc("/products", controllers.ProductsController)
	http.HandleFunc("/change_password", controllers.ChangePasswordController)
	http.HandleFunc("/forgot_password", controllers.ForgotPasswordController)
	// http.Handle("/something", middleware.Guard{})
}
