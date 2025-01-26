package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sqli/controllers"
	"sqli/middleware"
	"time"
)

// Generate CSRF Token
func generateCSRFToken() string {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		panic("Failed to generate CSRF token")
	}
	return base64.URLEncoding.EncodeToString(token)
}

// Set CSRF Token as HTTP-only Cookie
func setCSRFToken(w http.ResponseWriter) {
	token := generateCSRFToken()

	// Set the CSRF token in a secure, HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true, // Cannot be accessed via JavaScript
		Secure:   true, // Use HTTPS
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour), // Expiry date for token
	})
}

// Validate CSRF Token from Request
func validateCSRFToken(r *http.Request) bool {
	// Get CSRF token from the cookie
	cookie, err := r.Cookie("csrf_token")
	if err != nil {
		return false
	}

	// Check the CSRF token sent in the header (for non-form-based requests like AJAX)
	headerToken := r.Header.Get("X-CSRF-Token")
	if headerToken == "" {
		// For form submissions (POST, PUT, DELETE), the CSRF token is validated from the cookie
		return cookie != nil
	}

	// If there's a token in the header, compare it with the one from the cookie
	return headerToken == cookie.Value
}

// CSRF Protection Middleware
func csrfProtection(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Skip validation for safe methods like GET
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
			if !validateCSRFToken(r) {
				http.Error(w, "Forbidden: CSRF token missing or invalid", http.StatusForbidden)
				return
			}
		}
		next(w, r)
	}
}

// Set Security Headers for each response
func setSecurityHeaders(w http.ResponseWriter) {
	// Anti-clickjacking
	w.Header().Set("X-Frame-Options", "DENY")

	// X-Content-Type-Options
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Content Security Policy (CSP)
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' https://cdn.jsdelivr.net; style-src 'self' https://cdn.jsdelivr.net; img-src 'self'; connect-src 'self';")

	// Permissions Policy
	w.Header().Set("Permissions-Policy", "geolocation=(), microphone=()")

	// Cache Control
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// Set CSRF Token Header (for form-based requests)
	// This step can be omitted if you prefer to set the CSRF token in cookies as an HTTP-only cookie
	// w.Header().Set("X-CSRF-Token", generateCSRFToken())
}

// Router sets up the routes with proper security headers, CSRF protection, and controllers.
func Router() {
	// Apply CSRF and security headers to each route
	http.HandleFunc("/login", applySecurityHeaders(csrfProtection(controllers.LoginController)))
	http.HandleFunc("/", applySecurityHeaders(csrfProtection(controllers.LoginController)))
	http.HandleFunc("/products", applySecurityHeaders(csrfProtection(middleware.Guard(controllers.ProductsController))))
	http.HandleFunc("/admin", applySecurityHeaders(csrfProtection(middleware.Guard(controllers.AdminController))))
	http.HandleFunc("/change_password", applySecurityHeaders(csrfProtection(middleware.Guard(controllers.ChangePasswordController))))
	http.HandleFunc("/forgot_password", applySecurityHeaders(csrfProtection(controllers.ForgotPasswordController)))
	http.HandleFunc("/logout", applySecurityHeaders(csrfProtection(middleware.Guard(controllers.LogoutController))))
	// Add more routes as needed
}

// applySecurityHeaders wraps a handler to apply security headers and CSRF protection.
func applySecurityHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set security headers
		setSecurityHeaders(w)

		// Set CSRF token for GET requests
		if r.Method == http.MethodGet {
			setCSRFToken(w)
		}

		// Call the original handler
		handler(w, r)
	}
}
