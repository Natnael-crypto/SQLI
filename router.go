package main

import (
	"net/http"
	"sqli/controllers"
)

func Router() {
	http.HandleFunc("/vuln", controllers.VulnController)
	http.HandleFunc("/login", controllers.LoginController)
}
