package domain

import (
	"mpbe/errs"
	"mpbe/logger"

	"gorm.io/gorm"
)

type EmployeeRepositoryDB struct {
	db *gorm.DB
}

func NewEmployeeRepositoryDB(client *gorm.DB) EmployeeRepositoryDB {
	return EmployeeRepositoryDB{client}
}

func (e *EmployeeRepositoryDB) FindAll() ([]Employees, *errs.AppErr) {
	var err error
	var employees []Employees
	err = e.db.Find(&employees).Error
	if err != nil {
		logger.Error("error fetch data to employees table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error")
	}

	return employees, nil
}

func (e EmployeeRepositoryDB) FindByID(id int) (Employees, *errs.AppErr) {

	var employees Employees
	query := e.db.First(&employees, "employee_id = ?", id)
	if query != nil {
		logger.Error("error fetch data to employees table")
		return employees, errs.NewUnexpectedError("unexpected error")
	}
	return employees, nil
}
