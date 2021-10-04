package repository

import (
	"database/sql"
)

type Expenses interface {
}

type Category interface {
}

type Repository struct {
	Expenses
	Category
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
