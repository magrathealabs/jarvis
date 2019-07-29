package services

import "github.com/magrathealabs/jarvis/domain/repositories"

// AdminService class
type AdminService struct {
	AdminRepository repositories.AdminRepository
}

// DoSomething with admins
func (service *AdminService) DoSomething() error {
	return nil
}
