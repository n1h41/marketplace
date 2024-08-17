package repositories

import (
	"database/sql"

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
	var category entity.Category
	query := "select id, name, is_sub_category, parent_id from category"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Name, &category.IsSubCategory, &category.Parent)
		if err != nil {
			return nil, err
		}
		categoryList = append(categoryList, category)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categoryList, nil
}

func (p *productRepo) CreateCategory(param dto.CreateNewCategory) error {
	var query string
	var err error
	if param.IsSubCategory {
		query = "insert into category(name, is_sub_category, parent_id) values($1, $2, $3)"
		_, err = p.db.Exec(query, param.Name, param.IsSubCategory, param.ParentId)
		if err != nil {
			return err
		}
		return nil
	}
	query = "insert into category(name, is_sub_category) values($1, $2)"
	_, err = p.db.Exec(query, param.Name, param.IsSubCategory)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) CreateProduct(dto.AddProductModel) error {
	// query := "select id from category where name = $1"
	panic("unimplemented")
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{db: db}
}
