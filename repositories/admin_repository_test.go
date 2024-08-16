package repositories

import (
	"testing"

	"n1h41/marketplace/dto"
	"n1h41/marketplace/entity"
)

func createUser(t *testing.T) {
	user := entity.User{
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

	t.Log(r)
}

func TestLogin(t *testing.T) {
	createUser(t)

	params := dto.AdminLoginModel{
		Email:    "test@gmail.com",
		Password: "test1234",
	}

	err := repo.Login(params)
	if err != nil {
		t.Fatalf("Failed: %s", err)
	}
}

func TestCreateCategory(t *testing.T) {
	newCategory := dto.CreateNewCategory{
		Name:          "Food and Beverages",
		IsSubCategory: false,
	}
	err := repo.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	newCategory = dto.CreateNewCategory{
		Name:          "Pepsi",
		IsSubCategory: true,
		Parent:        "Food and Beverages",
	}

	err = repo.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
}
