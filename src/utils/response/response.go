package response

import (
	"encoding/json"
	"net/http"

	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

// RespondJSON write json response
func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

// RespondError write json response if error occurs
func RespondError(w http.ResponseWriter, err rerr.RestError) {
	RespondJSON(w, err.Status(), err)
}
