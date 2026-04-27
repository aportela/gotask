package middlewares

import (
	"net/http"
)

func isAuthenticated() bool {
	return true
}

func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"unauthorized","message":"missing authenticated session"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
