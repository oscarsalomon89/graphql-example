package repository

import (
	"fmt"
	"main/internal/models"
	"time"

	"github.com/jinzhu/gorm"
)

type EmployeeRepository interface {
	AddEmployee(employee *models.Employee) error
	GetEmployee(id int) (models.Employee, error)
	DeleteEmployee(id int) (models.Employee, error)
	GetEmployeesByRol(id int) ([]models.Employee, error)
	GetEmployees() ([]models.Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (repo *employeeRepository) GetEmployees() ([]models.Employee, error) {
	users := []models.Employee{}
	repo.db.Find(&users)

	return users, nil
}

func (repo *employeeRepository) GetEmployeesByRol(id int) ([]models.Employee, error) {
	users := []models.Employee{}
	repo.db.Where("idRol = ?", id).Find(&users)

	return users, nil
}

func (repo *employeeRepository) GetEmployee(id int) (models.Employee, error) {
	emp := models.Employee{}
	repo.db.Find(&emp, id)
	if emp.ID == 0 {
		fmt.Printf("No employee found with ID: %d\n", id)
		return emp, fmt.Errorf("No employee found with ID: %d", id)
	}
	return emp, nil
}

func (repo *employeeRepository) DeleteEmployee(id int) (models.Employee, error) {
	emp := models.Employee{}
	repo.db.Delete(&emp, id)

	if emp.ID == 0 {
		fmt.Printf("No employee found with ID: %d\n", id)
		return emp, fmt.Errorf("No employee found with ID: %d", id)
	}
	return emp, nil
}

func (repo *employeeRepository) AddEmployee(employee *models.Employee) error {
	role := models.Role{}
	repo.db.Find(&role, employee.IdRol)
	if role.ID == 0 {
		fmt.Printf("No role found with ID: %d\n", employee.IdRol)
		return fmt.Errorf("No role found with ID: %d", employee.IdRol)
	}

	createdAt := time.Now()
	employee.DateCreated = createdAt
	employee.DateUpdated = createdAt

	err := repo.db.Create(&employee).Error
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		return err
	}
	return nil
}
