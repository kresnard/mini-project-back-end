package service

import (
	"mpbe/domain"
	"mpbe/dto"
	"mpbe/errs"
	"mpbe/input"
	setupdb "mpbe/setupDB"
)

type EmployeeService interface {
	GetAllEmployee(dto.Pagination, int) (dto.Pagination, *errs.AppErr)
	GetEmployeeID(int) (domain.Employees, *errs.AppErr)
	CreateEmployee(input.EmployeeInput) (domain.Employees, *errs.AppErr)
	DltEmployee(int) (domain.Employees, *errs.AppErr)
	UpdtEmployee(int, input.EmployeeInput) (domain.Employees, *errs.AppErr)
}

type DefaultEmployeeService struct {
	repo domain.EmployeeRepository
}

func NewEmployeeService(repository domain.EmployeeRepository) DefaultEmployeeService {
	return DefaultEmployeeService{repository}
}

func (e DefaultEmployeeService) GetAllEmployee(p dto.Pagination, id int) (dto.Pagination, *errs.AppErr) {

	db := setupdb.GetClientDB()
	userRepositoryDB := domain.NewUserRepositoryDB(db)
	user, _ := userRepositoryDB.GetUserByID(id)
	if user.Email == "" {
		return p, errs.NewNotFoundError("user not found")
	}
	employees, err := e.repo.FindAll(p)
	if err != nil {
		return employees, err
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

func (e DefaultEmployeeService) DltEmployee(id int) (domain.Employees, *errs.AppErr) {
	employees, err := e.repo.DeleteEmployee(id)
	if err != nil {
		return employees, err
	}
	return employees, nil
}

func (e DefaultEmployeeService) UpdtEmployee(id int, input input.EmployeeInput) (domain.Employees, *errs.AppErr) {
	employee := domain.Employees{}
	employee.Name = input.Name
	employee.Gender = input.Gender
	employee.Birthday = input.Birthday
	employee.Department = input.Department
	employee.Status = input.Status

	employees, err := e.repo.UpdateEmployee(id, employee)
	if err != nil {
		return employees, err
	}
	return employees, nil
}