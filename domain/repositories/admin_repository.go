package repositories

import "github.com/magrathealabs/jarvis/domain/models"

// AdminRepository interface
type AdminRepository interface {
	All() ([]*models.Admin, error)
	Find(id int) (*models.Admin, error)
	Create(*models.Admin) (*models.Admin, error)
	Update(*models.Admin) (*models.Admin, error)
	Delete(*models.Admin) (*models.Admin, error)
}
