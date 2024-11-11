package main

import (
	"log"
	"net/http"
	"sqli/initializers"
	// "sqli/controllers"
	// "sqli/views"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.MigrateDB()
}

func main() {
	Router()

	address := "0.0.0.0:5001"
	log.Printf("Listening on %v\n", address)
	http.ListenAndServe(address, nil)
}
