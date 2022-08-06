package service

import (
	"fmt"
	"mpbe/domain"
	"mpbe/errs"
	"mpbe/input"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	user.Role = input.Role
	user.CreatedOn = time.Now()

	hashPassword, errBc := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errBc != nil {
		return user, errs.NewUnexpectedError("unexpected error")
	}
	fmt.Println(hashPassword)

	user.Password = string(hashPassword)

	users, err := u.repo.RegisterUserInput(user)
	if err != nil {
		return users, err
	}

	return users, nil
}
