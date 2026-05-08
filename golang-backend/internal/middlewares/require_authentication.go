package middlewares

import (
	"context"
	"net/http"
)

func isAuthenticated(ctx context.Context) bool {
	userId, success := GetUserIDFromContext(ctx)
	return success && userId != ""
}

func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(r.Context()) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"unauthorized","message":"missing authenticated session"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
