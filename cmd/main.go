package main

import (
	"log"
	"net/http"
	static "sqli"
	"sqli/initializers"
	// "sqli/controllers"
	// "sqli/views"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.MigrateDB()
	initializers.ParseTemplates(static.Templates)
}

func main() {
	Router()

	address := "0.0.0.0:5000"
	log.Printf("Listening on %v\n", address)
	http.ListenAndServe(address, nil)
}
