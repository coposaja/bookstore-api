package users

import (
	"strings"
	"time"

	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

// User struct representing User
type User struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`
	Status      string    `json:"status"`
	Passowrd    string    `json:"password"`
}

// Users is a slice of User
type Users []User

// PublicUser struct representing another User that not him
type PublicUser struct {
	ID          int       `json:"id"`
	DateCreated time.Time `json:"dateCreated"`
	Status      string    `json:"status"`
}

// PrivateUser struct representing himself
type PrivateUser struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`
	Status      string    `json:"status"`
}

// Validate method to validate User struct
func (user *User) Validate() rerr.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rerr.NewBadRequestError("Invalid Email")
	}

	user.Passowrd = strings.TrimSpace(user.Passowrd)
	if user.Passowrd == "" {
		return rerr.NewBadRequestError("Invalid Passowrd")
	}
	return nil
}
