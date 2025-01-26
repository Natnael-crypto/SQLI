package main

import (
	"log"
	"net/http"
	"sqli/initializers"
	"time"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.MigrateDB()
	initializers.ParseTemplates(static.Templates)
}

func setSecurityHeaders(w http.ResponseWriter) {
	// Anti-clickjacking
	w.Header().Set("X-Frame-Options", "DENY")

	// X-Content-Type-Options
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Content Security Policy
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self'; connect-src 'self';")

	// Permissions Policy
	w.Header().Set("Permissions-Policy", "geolocation=(), microphone=()")

	// Cache Control (for non-storable content)
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	// Anti-CSRF Token (ensure it's in your form)
	// w.Header().Set("X-CSRF-Token", "<CSRF_TOKEN>") // Replace with actual token generation logic
}

func main() {
	// Setup middleware to apply security headers
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Apply security headers
		setSecurityHeaders(w)

		// Your handler logic here
	})

	address := "0.0.0.0:5000"
	log.Printf("Listening on %v\n", address)

	// Create a custom server with timeouts
	server := &http.Server{
		Addr:              address,
		Handler:           nil, // Use default HTTP handler
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Handle errors from ListenAndServe
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
