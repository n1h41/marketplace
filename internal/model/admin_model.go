package model

import "mime/multipart"

type LoginAdminUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddCategoryReqeust struct {
	CategoryName  string `form:"category_name"`
	IsSubCategory bool   `form:"has_parent_category"`
	ParentId      string `form:"parent_id"`
}

type AddProductRequest struct {
	ProductName       string                  `json:"prodcutName"`
	ProductQuantity   int                     `json:"productQuantity"`
	ProductPrice      float32                 `json:"productPrice"`
	ProductImageFiles []*multipart.FileHeader `json:"productImageFiles"`
}
