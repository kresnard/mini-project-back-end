package service

import (
	"mpbe/domain"
	"mpbe/errs"
	"mpbe/input"
)

type UserService interface {
	CreateUser(input.UserInput) (domain.Users, *errs.AppErr)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}

func (u DefaultUserService) CreateUser(input input.UserInput) (domain.Users, *errs.AppErr) {
	user := domain.Users{}
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role

	users, err := u.repo.CreatUserInput(user)
	if err != nil {
		return users, err
	}
	return users, nil
}
