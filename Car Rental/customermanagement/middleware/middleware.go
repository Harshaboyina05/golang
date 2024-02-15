package middleware

import (
	"net/http"
)

// CORSMiddleware sets CORS headers
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001") // Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			// Preflight request, respond with success
			w.WriteHeader(http.StatusOK)
			return
		}

		// Serve the next handler
		next.ServeHTTP(w, r)
	})
}
