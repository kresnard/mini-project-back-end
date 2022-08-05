package service

import (
	"mpbe/domain"
	"mpbe/errs"
	"mpbe/input"
)

type EmployeeService interface {
	GetAllEmployee() ([]domain.Employees, *errs.AppErr)
	GetEmployeeID(int) (domain.Employees, *errs.AppErr)
	CreateEmployee(input.EmployeeInput) (domain.Employees, *errs.AppErr)
}

type DefaultEmployeeService struct {
	repo domain.EmployeeRepository
}

func NewEmployeeService(repository domain.EmployeeRepository) DefaultEmployeeService {
	return DefaultEmployeeService{repository}
}

func (e DefaultEmployeeService) GetAllEmployee() ([]domain.Employees, *errs.AppErr) {

	employees, err := e.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (e DefaultEmployeeService) GetEmployeeID(id int) (domain.Employees, *errs.AppErr) {
	employees, err := e.repo.FindByID(id)
	if err != nil {
		return employees, err
	}
	return employees, nil
}

func (e DefaultEmployeeService) CreateEmployee(input input.EmployeeInput) (domain.Employees, *errs.AppErr) {
	employee := domain.Employees{}
	employee.Name = input.Name
	employee.Gender = input.Gender
	employee.Birthday = input.Birthday
	employee.Department = input.Department
	employee.Status = input.Status

	employees, err := e.repo.CreateEmployeeInput(employee)
	if err != nil {
		return employees, err
	}
	return employees, nil
}
