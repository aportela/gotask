package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJSONError(w http.ResponseWriter, status int, code, message, debug string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errResp := map[string]interface{}{
		"APIError": true,
		"code":     code,
		"message":  message,
		"debug":    debug,
	}
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}
