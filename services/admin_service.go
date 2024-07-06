package services

import (
	"n1h41/marketplace/dto"
	"n1h41/marketplace/repositories"
)

type AdminService interface {
	Login(dto.AdminLoginModel) error
}

type adminService struct {
	repo repositories.AdminRepository
}

func AdminServiceConstructor(repo repositories.AdminRepository) AdminService {
	return &adminService{
		repo: repo,
	}
}

func (s *adminService) Login(data dto.AdminLoginModel) (err error) {
	err = s.repo.Login(data)
	return
}
