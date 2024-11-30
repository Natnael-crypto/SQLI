package main

import (
	"net/http"

	"sqli/controllers"
	"sqli/middleware"
)

func Router() {
	http.HandleFunc("/login", controllers.LoginController)
	http.HandleFunc("/", controllers.LoginController)
	http.HandleFunc("/products", middleware.Guard(controllers.ProductsController))
	http.HandleFunc("/admin", middleware.Guard(controllers.AdminController))
	http.HandleFunc("/change_password", middleware.Guard(controllers.ChangePasswordController))
	http.HandleFunc("/forgot_password", controllers.ForgotPasswordController)
	// http.Handle("/something", middleware.Guard{})
}
