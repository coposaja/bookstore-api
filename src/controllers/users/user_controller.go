package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/coposaja/bookstore-api/src/domains/users"
	"github.com/coposaja/bookstore-api/src/services"
	"github.com/coposaja/bookstore-api/src/utils/rerr"
	"github.com/coposaja/bookstore-api/src/utils/response"
	"github.com/gorilla/mux"
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

// GetUser handler to create a User
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(mux.Vars(r)["userId"], 10, 64)
	if err != nil {
		restErr := rerr.NewBadRequestError("UserId should be a number")
		response.RespondError(w, restErr)
		return
	}

	user, getErr := services.UserService.GetUser(int(userID))
	if getErr != nil {
		response.RespondError(w, getErr)
		return
	}

	response.RespondJSON(w, http.StatusOK, user)
}

// UpdateUser handler to update User data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		restErr := rerr.NewBadRequestError("Invalid JSON body")
		response.RespondError(w, restErr)
		return
	}

	userID, err := strconv.ParseInt(mux.Vars(r)["userId"], 10, 64)
	if err != nil {
		restErr := rerr.NewBadRequestError("UserId should be a number")
		response.RespondError(w, restErr)
		return
	}

	result, updateErr := services.UserService.UpdateUser(user, int(userID))
	if updateErr != nil {
		response.RespondError(w, updateErr)
		return
	}

	response.RespondJSON(w, http.StatusCreated, result)
}

// DeleteUser handler to create a User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(mux.Vars(r)["userId"], 10, 64)
	if err != nil {
		restErr := rerr.NewBadRequestError("UserId should be a number")
		response.RespondError(w, restErr)
		return
	}

	deleteErr := services.UserService.DeleteUser(int(userID))
	if deleteErr != nil {
		response.RespondError(w, deleteErr)
		return
	}

	response.RespondJSON(w, http.StatusOK, nil)
}
