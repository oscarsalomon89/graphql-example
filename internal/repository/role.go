package repository

import (
	"fmt"
	"main/internal/models"

	"github.com/jinzhu/gorm"
)

type RoleRepository interface {
	AddRole(repo *models.Role) error
	GetRoles() ([]models.Role, error)
	GetRole(id int) (models.Role, error)
	DeleteRole(id int) (models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (repo *roleRepository) GetRole(id int) (models.Role, error) {
	rol := models.Role{}
	repo.db.Find(&rol, id)
	if rol.ID == 0 {
		fmt.Printf("No employee found with ID: %d\n", id)
		return rol, fmt.Errorf("No employee found with ID: %d", id)
	}
	return rol, nil
}

func (repo *roleRepository) DeleteRole(id int) (models.Role, error) {
	rol := models.Role{}
	repo.db.Delete(&rol, id)

	if rol.ID == 0 {
		fmt.Printf("No role found with ID: %d\n", id)
		return rol, fmt.Errorf("No role found with ID: %d", id)
	}
	return rol, nil
}

func (repo *roleRepository) GetRoles() ([]models.Role, error) {
	roles := []models.Role{}
	repo.db.Find(&roles)

	return roles, nil
}

func (repo *roleRepository) AddRole(role *models.Role) error {
	err := repo.db.Create(&role).Error
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		return err
	}
	return nil
}
