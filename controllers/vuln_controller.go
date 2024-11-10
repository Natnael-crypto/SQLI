package controllers

import (
	"fmt"
	"net/http"
)


func VulnController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello this kinda works")
}