package repositories

import (
	"database/sql"
	"errors"

	"n1h41/marketplace/dto"
	"n1h41/marketplace/entity"
)

type AdminRepository interface {
	Login(dto.AdminLoginModel) error
	CreateProduct(dto.AddProductModel) error
	CreateCategory(dto.CreateNewCategory) error
}

type adminRepository struct {
	db *sql.DB
}

func (r *adminRepository) getParentCategoryId(categoryName string) int {
	var category entity.Category
	query := "select * from categories where name = $1"
	row, _ := r.db.Query(query, categoryName)
	err := row.Scan(category)
	if err != nil {
		return 0
	}
	return category.Id
}

func (r *adminRepository) CreateCategory(param dto.CreateNewCategory) error {
	var parentCategoryId int
	var query string
	var err error
	if param.IsSubCategory {
		parentCategoryId = r.getParentCategoryId(param.Parent)
		query = "insert into categories(name, is_sub_category, parent) values($1, $2, $3)"
		_, err = r.db.Exec(query, param.Name, param.IsSubCategory, parentCategoryId)
		if err != nil {
			return err
		}
		return nil
	}
	query = "insert into categories(name, is_sub_category) values($1, $2)"
	_, err = r.db.Exec(query, param.Name, param.IsSubCategory)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) CreateProduct(dto.AddProductModel) error {
	// query := "select id from categories where name = $1"
	panic("unimplemented")
}

func (r adminRepository) Login(params dto.AdminLoginModel) (err error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE email = $1 AND password = $2 AND role = 'Admin'", params.Email, params.Password)
	if row.Err() != nil {
		err = row.Err()
		return
	}
	if row.Scan() == sql.ErrNoRows {
		err = errors.New("User not found")
		return
	}
	return nil
}

func AdminRepoConstructor(db *sql.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}
