package services

import (
	"github.com/coposaja/bookstore-api/src/domains/users"
	"github.com/coposaja/bookstore-api/src/utils/crypto"
	"github.com/coposaja/bookstore-api/src/utils/date"
	"github.com/coposaja/bookstore-api/src/utils/rerr"
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, rerr.RestError)
	GetUser(int) (*users.User, rerr.RestError)
	UpdateUser(users.User, int) (*users.User, rerr.RestError)
	DeleteUser(int) rerr.RestError
	Search(string) ([]users.User, rerr.RestError)
}

func (s *userService) CreateUser(user users.User) (*users.User, rerr.RestError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	user.Passowrd, err = crypto.HashAndSalt(user.Passowrd)
	if err != nil {
		return nil, err
	}

	user.DateCreated = date.GetNow()
	user.Status = users.UserStatusActive

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

	curr.Passowrd, err = crypto.HashAndSalt(user.Passowrd)
	if err != nil {
		return nil, err
	}

	curr.FirstName = user.FirstName
	curr.LastName = user.LastName
	curr.Email = user.Email
	curr.Status = user.Status

	if err := curr.Validate(); err != nil {
		return nil, err
	}
	if err := curr.Update(); err != nil {
		return nil, err
	}

	return curr, nil
}

func (s *userService) DeleteUser(userID int) rerr.RestError {
	user := &users.User{ID: userID}
	return user.Delete()
}

func (s *userService) Search(status string) ([]users.User, rerr.RestError) {
	user := &users.User{}
	return user.Search(status)
}
