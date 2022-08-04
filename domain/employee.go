package domain

import "mpbe/errs"

type Employees struct {
	ID         string `json:"id" db:"employee_id" gorm:"column:employee_id"`
	Name       string `json:"name" db:"name" gorm:"column:name"`
	Gender     string `json:"gender" db:"gender" gorm:"column:gender"`
	Birthday   string `json:"birthday" db:"birthday" gorm:"column:birthday"`
	Department string `json:"department" db:"department" gorm:"column:department"`
	Status     string `json:"status" db:"status" gorm:"column:status"`
}

type EmployeeRepository interface {
	FindAll() ([]Employees, *errs.AppErr)
}
