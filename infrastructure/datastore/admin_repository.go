package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/magrathealabs/jarvis/domain/models"
	"github.com/magrathealabs/jarvis/domain/repositories"
)

// AdminRepository implementation
type AdminRepository struct {
	Conn *gorm.DB
}

// NewAdminRepository constructor
func NewAdminRepository(conn *gorm.DB) repositories.AdminRepository {
	return &AdminRepository{Conn: conn}
}

// All query
func (repository *AdminRepository) All() ([]*models.Admin, error) {
	var (
		admins []*models.Admin
		err    error
	)

	err = repository.Conn.Find(&admins).Error
	return admins, err
}

// Find query
func (repository *AdminRepository) Find(id int) (*models.Admin, error) {
	return nil, nil
}

// Create query
func (repository *AdminRepository) Create(*models.Admin) (*models.Admin, error) {
	return nil, nil
}

// Update query
func (repository *AdminRepository) Update(*models.Admin) (*models.Admin, error) {
	return nil, nil
}

// Delete query
func (repository *AdminRepository) Delete(*models.Admin) (*models.Admin, error) {
	return nil, nil
}
