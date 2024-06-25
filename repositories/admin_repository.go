package repositories

import (
	"database/sql"
	"errors"

	"n1h41/marketplace/models"
)

type AdminRepository interface {
	Login(models.AdminLoginModel) error
}

type adminRepository struct {
	db *sql.DB
}

func Constructor(db *sql.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r adminRepository) Login(params models.AdminLoginModel) (err error) {
	// TODO:
	row := r.db.QueryRow("select * from users where email = ? and password = ?", params.Email, params.Password)
	if row.Scan() == sql.ErrNoRows {
		err = errors.New("User not found")
		return
	}
	return nil
}
