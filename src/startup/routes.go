package startup

import (
	"net/http"

	"github.com/coposaja/bookstore-api/src/controllers/ping"
	"github.com/coposaja/bookstore-api/src/controllers/users"
)

func initRoutes() {
	router.HandleFunc("/ping", ping.Ping)

	router.HandleFunc("/users", users.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{userId}", users.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}", users.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{userId}", users.DeleteUser).Methods(http.MethodDelete)
}
