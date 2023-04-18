package usecase

import (
	"github.com/mujahxd/eventgenie/app/features/user"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository user.Repository
}

func NewService(repository user.Repository) user.Service {
	return &service{repository}
}

func (s *service) RegisterUser(input user.RegisterUserInput) (user.User, error) {

	user := user.User{}
	user.Name = input.Name
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
