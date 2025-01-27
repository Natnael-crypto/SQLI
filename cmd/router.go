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

	// X-Content-Type-Options to prevent MIME sniffing
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Content Security Policy (CSP) to mitigate XSS and other code injection attacks
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/; style-src 'self' https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/; img-src 'self'; connect-src 'self';")

	// Permissions Policy (formerly Feature Policy) to control access to certain browser features
	w.Header().Set("Permissions-Policy", "geolocation=(), microphone=()")

	// Cache Control: No caching to prevent stale or sensitive data from being cached
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// CSRF Token to protect against cross-site request forgery
	csrfToken := generateCSRFToken()

	// Set CSRF token as a cookie with SameSite attribute
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Path:     "/",
		HttpOnly: true,                    // Makes the cookie accessible only by the server (not JavaScript)
		Secure:   true,                    // Only send the cookie over HTTPS
		SameSite: http.SameSiteStrictMode, // Enforces SameSite cookie policy to prevent CSRF
	})

	// HTTP Strict Transport Security (HSTS) to enforce HTTPS connection (recommend max-age of 31536000 for production)
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

	// Cross-Site Scripting (XSS) Protection header
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Referrer Policy to control what information is sent in the Referer header
	w.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")
}

// Router sets up the routes with proper security headers, CSRF protection, and controllers.
func Router() {
	// Apply CSRF and security headers to each route
	http.HandleFunc("/login", applySecurityHeaders(csrfProtection(middleware.RateLimitMiddleware(controllers.LoginController))))
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
