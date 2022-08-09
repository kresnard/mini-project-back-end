package domain

import (
	"fmt"
	"mpbe/errs"
	"mpbe/logger"

	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(client *gorm.DB) UserRepositoryDB {
	return UserRepositoryDB{client}
}

func (u UserRepositoryDB) RegisterUserInput(users Users) (Users, *errs.AppErr) {
	err := u.db.Create(&users).Error

	if err != nil {
		logger.Error("error to registration user")
		return users, errs.NewUnexpectedError("Unexpected error")
	}
	return users, nil
}

func (u UserRepositoryDB) LoginUserInput(email string) (Users, *errs.AppErr) {

	var users Users
	err := u.db.First(&users, "email = ?", email).Error
	if err != nil {
		fmt.Println(email)
		logger.Error("error to login user")
		return users, errs.NewUnexpectedError("unexpected error")
	}
	return users, nil
}

func (u UserRepositoryDB) GetUserByID(id int) (Users, *errs.AppErr) {
	var users Users
	// err := u.db.First(&users, "user_id = ?", id).Error
	err := u.db.Where("user_id = ?", id).Find(&users).Error
	if err != nil {
		fmt.Println(id)
		logger.Error("error fetch data to users table")
		return users, errs.NewUnexpectedError("unexpected error")
	}
	return users, nil
}
