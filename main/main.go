package main

import (
	"net/http"
	// "sqli/controllers"
	// "sqli/views"
)

// import "sqli/vuln"

func main() {
	Router()
	http.ListenAndServe("0.0.0.0:5001", nil)
}

