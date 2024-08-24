package productrepo

import (
	"testing"

	"n1h41/marketplace/internal/model"
)

func TestCreateCategory(t *testing.T) {
	newCategory := model.AddCategoryReqeust{
		CategoryName:  "Food and Beverages",
		IsSubCategory: false,
	}
	err := productRepoMock.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	newCategory = model.AddCategoryReqeust{
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
