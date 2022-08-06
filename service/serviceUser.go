package service

import (
	"mpbe/domain"
	"mpbe/errs"
	"mpbe/input"
	"time"
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
	user.CreatedOn = time.Now()
	users, err := u.repo.RegisterUserInput(user)
	if err != nil {
		return users, err
	}

	// hashPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"code": 500,
	// 		"MSG":  "encryption error",
	// 	})
	// 	return
	// }

	return users, nil
}
