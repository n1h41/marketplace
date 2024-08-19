package dto

import "mime/multipart"

type AddProductModel struct {
	ProductName       string                  `json:"prodcutName"`
	ProductQuantity   int                     `json:"productQuantity"`
	ProductPrice      float32                 `json:"productPrice"`
	ProductImageFiles []*multipart.FileHeader `json:"productImageFiles"`
}
