package domain

import (
	"mpbe/dto"
	"mpbe/errs"
)

type Employees struct {
	ID         int    `json:"id" gorm:"primaryKey;column:employee_id"`
	Name       string `json:"name" gorm:"column:name"`
	Gender     string `json:"gender" gorm:"column:gender"`
	Birthday   string `json:"birthday" gorm:"column:birthday"`
	Department string `json:"department" gorm:"column:department"`
	Status     int    `json:"status" gorm:"column:status"`
}

type EmployeeRepository interface {
	FindAll(dto.Pagination) (dto.Pagination, *errs.AppErr)
	FindByID(int) (Employees, *errs.AppErr)
	CreateEmployeeInput(Employees) (Employees, *errs.AppErr)
	DeleteEmployee(int) (Employees, *errs.AppErr)
	UpdateEmployee(int, Employees) (Employees, *errs.AppErr)
}