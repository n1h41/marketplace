package repositories

import (
	"context"
	"database/sql"

	"github.com/georgysavva/scany/v2/sqlscan"

	"n1h41/marketplace/dto"
	"n1h41/marketplace/entity"
)

type ProductRepo interface {
	CreateProduct(dto.AddProductModel) error
	CreateCategory(dto.CreateNewCategory) error
	GetAllCategories() ([]entity.Category, error)
}

type productRepo struct {
	db *sql.DB
}

func (p *productRepo) GetAllCategories() ([]entity.Category, error) {
	var categoryList []entity.Category
	query := "select id, name, is_sub_category, parent_id from category"
	ctx := context.Background()
	if err := sqlscan.Select(ctx, p.db, &categoryList, query); err != nil {
		return nil, err
	}
	return categoryList, nil
}

func (p *productRepo) CreateCategory(param dto.CreateNewCategory) error {
	var query string
	var err error
	if param.IsSubCategory {
		query = "insert into category(name, is_sub_category, parent_id) values($1, $2, $3)"
		_, err = p.db.Exec(query, param.CategoryName, param.IsSubCategory, param.ParentId)
		if err != nil {
			return err
		}
		return nil
	}
	query = "insert into category(name, is_sub_category) values($1, $2)"
	_, err = p.db.Exec(query, param.CategoryName, param.IsSubCategory)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) CreateProduct(dto.AddProductModel) error {
	panic("unimplemented")
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{db: db}
}
