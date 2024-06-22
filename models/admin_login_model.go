package models

type AdminLoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
