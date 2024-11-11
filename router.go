package main

import (
	"net/http"

	"sqli/controllers"
	"sqli/middleware"
)

func Router() {
	http.HandleFunc("/login", controllers.LoginController)
	http.Handle("/something", middleware.Guard{})
}
