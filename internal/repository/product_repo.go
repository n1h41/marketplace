package repository

import (
	"context"
	"database/sql"

	"github.com/georgysavva/scany/v2/sqlscan"

	"n1h41/marketplace/internal/domain/productdmn"
	"n1h41/marketplace/internal/model"
)

type ProductRepo interface {
	CreateProduct(model.AddProductRequest) error
	CreateCategory(productdmn.Category) error
	GetAllCategories() ([]productdmn.Category, error)
}

type productRepo struct {
	db *sql.DB
}

func (p *productRepo) GetAllCategories() ([]productdmn.Category, error) {
	var categoryList []productdmn.Category
	query := "select id, name, is_sub_category, parent_id from category"
	ctx := context.Background()
	if err := sqlscan.Select(ctx, p.db, &categoryList, query); err != nil {
		return nil, err
	}
	return categoryList, nil
}

func (p *productRepo) CreateCategory(category productdmn.Category) error {
	var query string
	var err error
	if category.IsSubCategory {
		query = "insert into category(name, is_sub_category, parent_id) values($1, $2, $3)"
		_, err = p.db.Exec(query, category.Name, category.IsSubCategory, category.Parent)
		if err != nil {
			return err
		}
		return nil
	}
	query = "insert into category(name, is_sub_category) values($1, $2)"
	_, err = p.db.Exec(query, category.Name, category.IsSubCategory)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) CreateProduct(model.AddProductRequest) error {
	panic("unimplemented")
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{db: db}
}
