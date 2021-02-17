package services

import (
	"github.com/coposaja/bookstore-api/src/domains/users"
	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, rerr.RestError)
	GetUser(int) (*users.User, rerr.RestError)
	UpdateUser(users.User, int) (*users.User, rerr.RestError)
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

func (s *userService) GetUser(userID int) (*users.User, rerr.RestError) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(user users.User, userID int) (*users.User, rerr.RestError) {
	curr, err := s.GetUser(userID)
	if err != nil {
		return nil, err
	}

	curr.FirstName = user.FirstName
	curr.LastName = user.LastName
	curr.Email = user.Email

	if err := curr.Validate(); err != nil {
		return nil, err
	}
	if err := curr.Update(); err != nil {
		return nil, err
	}

	return curr, nil
}
