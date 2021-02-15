package startup

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
}
