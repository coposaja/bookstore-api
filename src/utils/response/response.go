package response

import (
	"encoding/json"
	"net/http"
)

// RespondJSON write json response
func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}
