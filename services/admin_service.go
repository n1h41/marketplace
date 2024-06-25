package services

import (
	"n1h41/marketplace/models"
	"n1h41/marketplace/repositories"
)

type AdminService interface {
	Login(models.AdminLoginModel) error
}

type adminService struct {
	repo repositories.AdminRepository
}

func Constructor(repo repositories.AdminRepository) AdminService {
	return &adminService{
		repo: repo,
	}
}

func (s *adminService) Login(data models.AdminLoginModel) (err error) {
	err = s.repo.Login(data)
	return
}
