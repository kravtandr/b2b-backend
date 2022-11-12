package models

type CategorieId struct {
	Id int64
}

type Category struct {
	Id   int64
	Name string
}

type Categories []Category
