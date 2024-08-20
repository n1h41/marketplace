package dto

type CreateNewCategory struct {
	CategoryName  string `db:"category_name"`
	IsSubCategory bool   `db:"has_parent_category"`
	ParentId      int    `db:"parent_id"`
}
