package main

import (
	"log"
	"net/http"
	"sqli/initializer"
	// "sqli/controllers"
	// "sqli/views"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
	initializer.MigrateDB()
}

func main() {
	Router()

	address := "0.0.0.0:5001"
	log.Printf("Listening on %v\n", address)
	http.ListenAndServe(address, nil)
}
