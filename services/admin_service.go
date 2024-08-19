package services

import (
	"n1h41/marketplace/dto"
	"n1h41/marketplace/repositories"
)

type AdminService interface {
	Login(dto.AdminLoginModel) error
	CreateProduct(dto.AddProductModel) error
}

type adminService struct {
	repo repositories.AdminRepo
}

func (s *adminService) Login(data dto.AdminLoginModel) (err error) {
	err = s.repo.Login(data)
	return
}

// CreateProduct implements AdminService.
func (s *adminService) CreateProduct(dto.AddProductModel) (err error) {
	panic("unimplemented")
}

func AdminServiceConstructor(repo repositories.AdminRepo) AdminService {
	return &adminService{
		repo: repo,
	}
}
