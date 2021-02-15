package services

import (
	"github.com/coposaja/bookstore-api/src/domains/users"
	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, rerr.RestError)
}

func (s *userService) CreateUser(user users.User) (*users.User, rerr.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
