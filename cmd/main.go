package main

import (
	"log"
	"net/http"
	static "sqli"
	"sqli/initializers"
	"time"
)

// Initialize the application
func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.MigrateDB()
	initializers.ParseTemplates(static.Templates)
}

// Function to set security headers
func setSecurityHeaders(w http.ResponseWriter) {
	// Anti-clickjacking
	w.Header().Set("X-Frame-Options", "DENY")

	// X-Content-Type-Options
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Content Security Policy
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self'; connect-src 'self';")

	// Permissions Policy
	w.Header().Set("Permissions-Policy", "geolocation=(), microphone=()")

	// Cache control for sensitive resources
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

// Function to set cache headers for cacheable resources
func setCacheHeaders(w http.ResponseWriter) {
	// Cache control for cacheable resources (e.g., static content like images, CSS, JS)
	w.Header().Set("Cache-Control", "public, max-age=86400, s-maxage=3600")         // Cache for 24 hours, shared caches for 1 hour
	w.Header().Set("Expires", time.Now().Add(24*time.Hour).Format(http.TimeFormat)) // Expires in 24 hours
}

// Main function to start the server
func main() {
	// Setup middleware to apply security and cache headers
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Apply security headers
		setSecurityHeaders(w)

		// Set cache headers based on resource type
		if req.URL.Path == "/favicon.ico" || req.URL.Path == "/robots.txt" || req.URL.Path == "/sitemap.xml" {
			setCacheHeaders(w) // Set cache for static resources like favicon.ico, robots.txt, etc.
		}

		// Your handler logic here
		// For example, render a page, serve static content, etc.
		http.ServeFile(w, req, "index.html") // Replace with your actual handler logic
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
