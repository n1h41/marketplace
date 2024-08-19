package dto

type CreateNewCategory struct {
	Name          string `db:"name"`
	IsSubCategory bool   `db:"is_sub_category"`
	ParentId      int    `db:"parent_id"`
}
