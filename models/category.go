package models

// Category fetch from db
type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
