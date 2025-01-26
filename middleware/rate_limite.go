package middleware

import (
	"net/http"
	"sync"
	"time"
)

var (
	rateLimitMap = make(map[string]*rateLimitInfo)
	mu           sync.Mutex
)

// Rate limit information for each IP address
type rateLimitInfo struct {
	requests  int
	resetTime time.Time
}

// RateLimitMiddleware is a middleware to enforce rate limiting
func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		mu.Lock()
		defer mu.Unlock()

		// Initialize rate limit info for the IP if it doesn't exist
		if _, exists := rateLimitMap[ip]; !exists {
			rateLimitMap[ip] = &rateLimitInfo{
				requests:  0,
				resetTime: time.Now().Add(time.Minute), // Reset limit every minute
			}
		}

		// Get rate limit info for the IP
		rateInfo := rateLimitMap[ip]

		// Reset the rate limit counter if the time period has passed
		if time.Now().After(rateInfo.resetTime) {
			rateInfo.requests = 0
			rateInfo.resetTime = time.Now().Add(time.Minute) // Reset in the next minute
		}

		// Check if the IP has exceeded the rate limit (e.g., 5 requests per minute)
		if rateInfo.requests >= 5 {
			http.Error(w, "Rate limit exceeded, please try again later.", http.StatusTooManyRequests)
			return
		}

		// Increment the request counter
		rateInfo.requests++

		// Call the next handler in the chain
		next(w, r)
	}
}
