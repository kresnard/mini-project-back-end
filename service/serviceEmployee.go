package service

import (
	"mpbe/domain"
	"mpbe/errs"
)

type EmployeeService interface {
	GetAllEmployee() ([]domain.Employees, *errs.AppErr)
}

type DefaultEmployeeService struct {
	repo domain.EmployeeRepository
}

func (s DefaultEmployeeService) GetAllEmployee() ([]domain.Employees, *errs.AppErr) {

	employees, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func NewEmployeeService(repository domain.EmployeeRepository) DefaultEmployeeService {
	return DefaultEmployeeService{repository}
}
