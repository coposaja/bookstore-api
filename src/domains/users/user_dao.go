package users

import (
	"fmt"

	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

var (
	userDB = make(map[int]*User)
)

// Save User to DB
func (user *User) Save() rerr.RestError {
	for _, v := range userDB {
		if v.Email == user.Email {
			return rerr.NewBadRequestError(fmt.Sprintf("User with email %s already exists", user.Email))
		}
	}

	userDB[user.ID] = user
	return nil
}

// Get User from DB by UserID
func (user *User) Get() rerr.RestError {
	current := userDB[user.ID]
	if current == nil {
		return rerr.NewNotFoundError(fmt.Sprintf("User with id %d is not found", user.ID))
	}

	user.FirstName = current.FirstName
	user.LastName = current.LastName
	user.Email = current.Email
	user.DateCreated = current.DateCreated

	return nil
}
