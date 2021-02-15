package ping

import (
	"net/http"

	"github.com/coposaja/bookstore-api/src/utils/response"
)

// Ping handler to test connection
func Ping(w http.ResponseWriter, r *http.Request) {
	response.RespondJSON(w, http.StatusOK, "Pong!")
}
