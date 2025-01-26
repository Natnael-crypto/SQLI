package main

import (
	"log"
	"net/http"
	static "sqli"
	"sqli/initializers"

	// "sqli/controllers"
	// "sqli/views"
	"time"
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

	// Create an http.Server with timeouts
	server := &http.Server{
		Addr:              address,
		Handler:           nil,              // Use default HTTP handler
		ReadTimeout:       10 * time.Second, // Maximum duration for reading the entire request, including the body
		WriteTimeout:      10 * time.Second, // Maximum duration before timing out writes of the response
		IdleTimeout:       15 * time.Second, // Maximum amount of time to wait for the next request when keep-alives are enabled
		ReadHeaderTimeout: 5 * time.Second,  // Timeout for reading request headers
	}

	// Handle errors from ListenAndServe
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
