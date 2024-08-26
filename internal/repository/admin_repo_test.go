package repository

import (
	"testing"

	"n1h41/marketplace/internal/domain/coredmn"
	"n1h41/marketplace/internal/model"
)

func createUser(t *testing.T) {
	user := coredmn.User{
		FirstName: "Nihal",
		LastName:  "Abdulla",
		Email:     "test@gmail.com",
		Password:  "test1234",
		Role:      "Admin",
	}

	r, err := testDB.Exec("INSERT INTO users(email, password, first_name, last_name, role) values($1, $2, $3, $4, $5)", user.Email, user.Password, user.FirstName, user.LastName, user.Role)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	t.Log(r.RowsAffected())
}

func TestLogin(t *testing.T) {
	createUser(t)

	params := model.LoginAdminUserRequest{
		Email:    "test@gmail.com",
		Password: "test1234",
	}

	err := adminRepoMock.Login(params)
	if err != nil {
		t.Fatalf("Failed: %s", err)
	}
}
