package domain

import (
	"mpbe/errs"
	"time"
)

type Users struct {
	ID        int       `json:"user_id" gorm:"column:user_id"`
	Username  string    `json:"username" gorm:"column:username"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	Role      string    `json:"role" gorm:"column:role"`
	CreatedOn time.Time `json:"created_on" gorm:"column:created_on"`
}

type UserRepository interface {
	RegisterUserInput(Users) (Users, *errs.AppErr)
}
