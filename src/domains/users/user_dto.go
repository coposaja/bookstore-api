package users

import (
	"strings"

	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

// User struct representing User
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}

// Validate method to validate User struct
func (user *User) Validate() rerr.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rerr.NewBadRequestError("Invalid Email")
	}
	return nil
}
