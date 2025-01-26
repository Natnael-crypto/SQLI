package main

import (
	"log"
	"net/http"
	static "sqli"
	"sqli/initializers"
	"os"
	"time"
)

// Initializes environment, DB, migrations, and templates.
func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.MigrateDB()
	initializers.ParseTemplates(static.Templates)
}

// Sets security headers for all responses.
func setSecurityHeaders(w http.ResponseWriter) {
	// Anti-clickjacking
	w.Header().Set("X-Frame-Options", "DENY")

	// X-Content-Type-Options
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Content Security Policy
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self'; connect-src 'self';")

	// Permissions Policy
	w.Header().Set("Permissions-Policy", "geolocation=(), microphone=()")

	// Cache Control for sensitive content
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

// Sets non-cacheable headers for static resources (like favicon, robots.txt, sitemap.xml)
func setNonCacheableHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

// Serve the requested file or return a 404 if the file does not exist.
func serveStaticFile(w http.ResponseWriter, req *http.Request, filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, req) // Return a 404 if the file does not exist
		return
	}

	http.ServeFile(w, req, filePath)
}

func main() {
	// Static file handlers with cache control headers
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
		setNonCacheableHeaders(w)
		// Serve the favicon file
		serveStaticFile(w, req, "static/favicon.ico")
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, req *http.Request) {
		setNonCacheableHeaders(w)
		// Serve the robots.txt file or return 404 if not found
		serveStaticFile(w, req, "static/robots.txt")
	})

	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, req *http.Request) {
		setNonCacheableHeaders(w)
		// Serve the sitemap.xml file or return 404 if not found
		serveStaticFile(w, req, "static/sitemap.xml")
	})

	// Setup middleware to apply security headers to all routes
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Apply general security headers
		setSecurityHeaders(w)

		// Your main route logic here (for example, serving HTML pages or API responses)
		w.Write([]byte("Hello, World! This is your secure server."))
	})

	// Define the server address
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
