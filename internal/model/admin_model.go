package model

import "mime/multipart"

type LoginAdminUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddCategoryReqeust struct {
	CategoryName  string `db:"category_name"`
	IsSubCategory bool   `db:"has_parent_category"`
	ParentId      int    `db:"parent_id"`
}

type AddProductRequest struct {
	ProductName       string                  `json:"prodcutName"`
	ProductQuantity   int                     `json:"productQuantity"`
	ProductPrice      float32                 `json:"productPrice"`
	ProductImageFiles []*multipart.FileHeader `json:"productImageFiles"`
}
