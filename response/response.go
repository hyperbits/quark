package response

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RespondAndLogError(w http.ResponseWriter, code int, message, details string) {
	log.WithFields(log.Fields{
		"status":  code,
		"details": details,
	}).Error(message)
	RespondWithError(w, code, message)
}
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
