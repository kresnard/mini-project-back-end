package domain

import (
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

func (u UserRepositoryDB) RegistrationUserInput(users Users) (Users, *errs.AppErr) {
	err := u.db.Create(&users)

	if err != nil {
		logger.Error("error to registrarion user")
		return users, errs.NewUnexpectedError("Unexpected error")
	}
	return users, nil
}
