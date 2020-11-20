package services

import (
	"main/internal/models"

	"main/internal/repository"
)

//go:generate mockgen -destination=../utils/test/mocks/ingate_files_service_mock.go -package=mocks -source=./ingate_files_service.go

type EmployeService interface {
	Save(employee *models.Employee) error
	Get(id int) (models.Employee, error)
	GetAll() ([]models.Employee, error)
}

type employeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) EmployeService {
	return &employeService{employeeRepository: employeeRepository}
}

func (svc *employeService) Get(id int) (models.Employee, error) {
	employee, err := svc.employeeRepository.GetEmployee(id)

	return employee, err
}

func (svc *employeService) GetAll() ([]models.Employee, error) {
	employees, err := svc.employeeRepository.GetEmployees()

	return employees, err
}

func (svc *employeService) Save(employee *models.Employee) error {
	err := svc.employeeRepository.AddEmployee(employee)

	return err
}
