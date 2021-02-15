package startup

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

// StartApp initialize routing and start web server
func StartApp() {
	initRoutes()
	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      router,
	}

	log.Fatal(srv.ListenAndServe())
}
