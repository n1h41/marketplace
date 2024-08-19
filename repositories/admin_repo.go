package repositories

import (
	"database/sql"
	"errors"

	"n1h41/marketplace/dto"
)

type AdminRepo interface {
	Login(dto.AdminLoginModel) error
}

type adminRepo struct {
	db *sql.DB
}

func (a adminRepo) Login(params dto.AdminLoginModel) (err error) {
	row := a.db.QueryRow("SELECT * FROM users WHERE email = $1 AND password = $2 AND role = 'Admin'", params.Email, params.Password)
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

func NewAdminRepo(db *sql.DB) AdminRepo {
	return &adminRepo{
		db: db,
	}
}
