package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON mengirimkan respons dengan format JSON yang konsisten.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]interface{}{
		"status": status,
		"data":   payload,
	}
	json.NewEncoder(w).Encode(response)
}

// RespondWithError mengirimkan respons error dengan pesan yang jelas.
func RespondWithError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]interface{}{
		"status": status,
		"error":  message,
	}
	json.NewEncoder(w).Encode(response)
}
