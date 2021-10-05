package repository

import (
	"database/sql"

	go_bot "github.com/nexeranet/go-bot"
)

type Expenses interface {
}

type Category interface {
	GetOne(string) (go_bot.Category, error)
}

type Repository struct {
	Expenses
	Category
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Expenses: NewExpensesSqlite3(db),
		Category: NewCategorySqlite3(db),
	}
}