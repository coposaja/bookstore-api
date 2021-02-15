package services

import "github.com/coposaja/bookstore-api/src/domains/users"

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, error)
}

func (s *userService) CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
