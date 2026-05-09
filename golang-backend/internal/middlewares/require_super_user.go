package middlewares

import (
	"log"
	"net/http"
)

func isSuperUser() bool {
	return true
}

func RequireSuperUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isSuperUser() {
			log.Printf("forbidden request to %s", r.URL.Path)
			writeJSONError(w, http.StatusForbidden,
				"REQUIRE_SUPER_USER_MIDDLEWARE_ERROR",
				"Forbidden",
				"")
			return
		}
		next.ServeHTTP(w, r)
	})
}
