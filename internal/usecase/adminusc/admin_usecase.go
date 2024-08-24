package adminusc

import (
	"n1h41/marketplace/internal/model"
	"n1h41/marketplace/internal/repository/adminrepo"
)

type AdminUsecase interface {
	Login(model.LoginAdminUserRequest) error
	CreateProduct(model.AddProductRequest) error
}

type adminUsecase struct {
	repo adminrepo.AdminRepo
}

func (a adminUsecase) Login(data model.LoginAdminUserRequest) (err error) {
	err = a.repo.Login(data)
	return
}

// CreateProduct implements AdminUsecase.
func (a adminUsecase) CreateProduct(model.AddProductRequest) (err error) {
	panic("unimplemented")
}

func NewAdminUsc(repo adminrepo.AdminRepo) AdminUsecase {
	return adminUsecase{
		repo: repo,
	}
}
