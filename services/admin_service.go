package services

import (
	"database/sql"

	"n1h41/marketplace/models"
)

type AdminService interface {
	Login(models.AdminLoginModel) error
}

type adminService struct {
	db *sql.DB
}

func Constructor(db *sql.DB) AdminService {
	return &adminService{
		db: db,
	}
}

func (s *adminService) Login(data models.AdminLoginModel) (err error) {
	rows, err := s.db.Query("SELECT * FROM admin WHERE email = ? AND password = ?", data.Email, data.Password)
	if err != nil {
		return
	}
	defer rows.Close()
	return nil
}
