package repository

import (
	"testing"

	"n1h41/marketplace/internal/domain/productdmn"
)

func TestCreateCategory(t *testing.T) {
	newCategory := productdmn.Category{
		Name:          "Food and Beverages",
		IsSubCategory: false,
	}
	err := productRepoMock.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	newCategory = productdmn.Category{
		Name:          "Pepsi",
		IsSubCategory: true,
		Parent:        1,
	}

	err = productRepoMock.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
}

func TestGetAllCategories(t *testing.T) {
	result, err := productRepoMock.GetAllCategories()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
	t.Log(len(result))
}
