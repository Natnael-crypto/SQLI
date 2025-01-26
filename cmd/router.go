package main

import (
	"net/http"
	"sqli/controllers"
	"sqli/middleware"
)

func setSecurityHeaders(w http.ResponseWriter) {
	// Anti-clickjacking
	w.Header().Set("X-Frame-Options", "DENY")

	// X-Content-Type-Options
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Content Security Policy (CSP)
	// Content Security Policy
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' https://maxcdn.bootstrapcdn.com; style-src 'self' https://maxcdn.bootstrapcdn.com; img-src 'self'; connect-src 'self';")

	// Permissions Policy
	w.Header().Set("Permissions-Policy", "geolocation=(), microphone=()")

	// Cache Control
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// Add Anti-CSRF Token Header (if you implement CSRF tokens)
	// w.Header().Set("X-CSRF-Token", generateCSRFToken())
}

// Router sets up the routes with proper security headers and controllers.
func Router() {
	// Apply security headers to each route
	http.HandleFunc("/login", applySecurityHeaders(controllers.LoginController))
	http.HandleFunc("/", applySecurityHeaders(controllers.LoginController))
	http.HandleFunc("/products", middleware.Guard(applySecurityHeaders(controllers.ProductsController)))
	http.HandleFunc("/admin", middleware.Guard(applySecurityHeaders(controllers.AdminController)))
	http.HandleFunc("/change_password", middleware.Guard(applySecurityHeaders(controllers.ChangePasswordController)))
	http.HandleFunc("/forgot_password", applySecurityHeaders(controllers.ForgotPasswordController))
	http.HandleFunc("/logout", middleware.Guard(applySecurityHeaders(controllers.LogoutController)))
	// Add more routes as needed
}

// applySecurityHeaders wraps a handler to apply security headers.
func applySecurityHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set security headers
		setSecurityHeaders(w)
		// Call the original handler
		handler(w, r)
	}
}
