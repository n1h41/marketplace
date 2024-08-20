package repositories

import (
	"testing"

	"n1h41/marketplace/dto"
)

func TestCreateCategory(t *testing.T) {
	newCategory := dto.CreateNewCategory{
		CategoryName:  "Food and Beverages",
		IsSubCategory: false,
	}
	err := productRepoMock.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	newCategory = dto.CreateNewCategory{
		CategoryName:  "Pepsi",
		IsSubCategory: true,
		ParentId:      1,
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
