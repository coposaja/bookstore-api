package users

import (
	"encoding/json"
	"net/http"

	"github.com/coposaja/bookstore-api/src/domains/users"
	"github.com/coposaja/bookstore-api/src/services"
	"github.com/coposaja/bookstore-api/src/utils/response"
)

// CreateUser handler to create a User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.RespondJSON(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		response.RespondJSON(w, http.StatusBadRequest, "Some errors")
		return
	}

	response.RespondJSON(w, http.StatusCreated, result)
}
