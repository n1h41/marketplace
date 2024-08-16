package dto

type CreateNewCategory struct {
	Name          string `db:"name"`
	IsSubCategory bool   `db:"is_sub_category"`
	Parent        string `db:"parent"`
}
