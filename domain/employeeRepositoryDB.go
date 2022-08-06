package domain

import (
	"fmt"
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
	err := e.db.First(&employees, "employee_id = ?", id)
	if err != nil {
		logger.Error("error fetch data to employees table")
		return employees, errs.NewUnexpectedError("unexpected error")
	}
	return employees, nil
}

func (e EmployeeRepositoryDB) CreateEmployeeInput(employees Employees) (Employees, *errs.AppErr) {

	err := e.db.Create(&employees).Error
	fmt.Println(err)

	if err != nil {
		// fmt.Println("err", err.Error())
		logger.Error("error to create employee data")
		return employees, errs.NewUnexpectedError("unexpected error")
	}
	return employees, nil
}

func (e EmployeeRepositoryDB) DeleteEmployee(id int) (Employees, *errs.AppErr) {

	var employees Employees
	err := e.db.Delete(&employees, "employee_id = ?", id)
	if err != nil {
		logger.Error("error to delete employee data")
		return employees, errs.NewUnexpectedError("unexpected error")
	}
	return employees, nil

}

func (e EmployeeRepositoryDB) UpdateEmployee(id int, employees Employees) (Employees, *errs.AppErr) {

	err := e.db.Model(&employees).Where("employee_id = ?", id).Updates(employees)

	if err != nil {
		logger.Error("error to update employee data")
		return employees, errs.NewUnexpectedError("unexpected error")
	}
	return employees, nil
}
