package repositories

import (
	"database/sql"
	"errors"

	"n1h41/marketplace/dto"
)

type AdminRepository interface {
	Login(dto.AdminLoginModel) error
}

type adminRepository struct {
	db *sql.DB
}

func AdminRepoConstructor(db *sql.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r adminRepository) Login(params dto.AdminLoginModel) (err error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE email = $1 AND password = $2", params.Email, params.Password)
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
