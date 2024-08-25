package adminusc

import (
	"n1h41/marketplace/internal/domain/productdmn"
	"n1h41/marketplace/internal/model"
	"n1h41/marketplace/internal/repository/adminrepo"
	"n1h41/marketplace/internal/repository/productrepo"
)

type AdminUsecase interface {
	Login(model.LoginAdminUserRequest) error
	CreateProduct(model.AddProductRequest) error
	CreateCategory(productdmn.Category) error
	ListCategories() ([]productdmn.Category, error)
}

type adminUsecase struct {
	adminRepo   adminrepo.AdminRepo
	productRepo productrepo.ProductRepo
}

func NewAdminUsc(adminRepo adminrepo.AdminRepo, productRepo productrepo.ProductRepo) AdminUsecase {
	return adminUsecase{
		adminRepo:   adminRepo,
		productRepo: productRepo,
	}
}

func (a adminUsecase) Login(data model.LoginAdminUserRequest) (err error) {
	err = a.adminRepo.Login(data)
	return
}

// CreateProduct implements AdminUsecase.
func (a adminUsecase) CreateProduct(model.AddProductRequest) (err error) {
	panic("unimplemented")
}

// CreateCategory implements AdminUsecase.
func (a adminUsecase) CreateCategory(newCategory productdmn.Category) error {
	err := a.productRepo.CreateCategory(newCategory)
	return err
}

// ListCategories implements AdminUsecase.
func (a adminUsecase) ListCategories() (categoryList []productdmn.Category, err error) {
	categoryList, err = a.productRepo.GetAllCategories()
	return
}
