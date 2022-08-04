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

func (s *EmployeeRepositoryDB) FindAll() ([]Employees, *errs.AppErr) {
	var err error
	var employees []Employees
	err = s.db.Find(&employees).Error
	if err != nil {
		logger.Error("error fetch data .to employees table " + err.Error())
		return nil, errs.NewUnexpectedError("test")
	}

	return employees, nil
}
