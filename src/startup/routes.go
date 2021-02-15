package startup

import (
	"net/http"

	"github.com/coposaja/bookstore-api/src/utils/response"
)

func initRoutes() {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response.RespondJSON(w, http.StatusOK, "pong")
	})
}
