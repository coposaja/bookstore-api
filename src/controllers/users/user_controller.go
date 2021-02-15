package users

import (
	"encoding/json"
	"net/http"

	"github.com/coposaja/bookstore-api/src/domains/users"
	"github.com/coposaja/bookstore-api/src/services"
	"github.com/coposaja/bookstore-api/src/utils/rerr"
	"github.com/coposaja/bookstore-api/src/utils/response"
)

// CreateUser handler to create a User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		restErr := rerr.NewBadRequestError("Invalid JSON body")
		response.RespondError(w, restErr)
		return
	}

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		response.RespondError(w, saveErr)
		return
	}

	response.RespondJSON(w, http.StatusCreated, result)
}
