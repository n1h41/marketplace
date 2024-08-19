package entity

type Category struct {
	Id            int    `db:"id"`
	Name          string `db:"name"`
	IsSubCategory bool   `db:"is_sub_category"`
	Parent        int    `db:"parent_id"`
}
