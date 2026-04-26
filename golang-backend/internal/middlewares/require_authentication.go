package middlewares

import (
	"encoding/json"
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
			resp := errorResponse{
				Message: "missing authenticated session",
			}
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}
