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
