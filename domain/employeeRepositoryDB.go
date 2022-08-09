package domain

import (
	"fmt"
	"mpbe/dto"
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

func (e *EmployeeRepositoryDB) FindAll(pagination dto.Pagination) (dto.Pagination, *errs.AppErr) {
	var p dto.Pagination
	tr := 0
	offset := pagination.Page * pagination.Limit

	var employees []Employees
	errFind := e.db.Limit(pagination.Limit).Offset(offset).Find(&employees).Error
	if errFind != nil {
		return p, nil
	}

	pagination.Rows = employees
	totalRows := int64(tr)

	errCount := e.db.Model(employees).Count(&totalRows).Error

	if errCount != nil {
		return p, errs.NewUnexpectedError("unexpected error")
	}

	return pagination, nil
}

func (e EmployeeRepositoryDB) FindByID(id int) (Employees, *errs.AppErr) {

	var employees Employees
	// err := e.db.First(&employees, "employee_id = ?", id).Error
	err := e.db.Where("employee_id = ?", id).Find(&employees).Error
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
