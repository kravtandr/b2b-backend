package models

import "database/sql"

type CategoryId struct {
	Id int64
}

type SearchProducts struct {
	name string
}

type Category struct {
	Id          int64
	Name        string
	Description sql.NullString
}

type Categories []Category
