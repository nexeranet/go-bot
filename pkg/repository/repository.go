package repository

import (
	"database/sql"

	go_bot "github.com/nexeranet/go-bot"
)

type Expenses interface {
	Create(category string, amount int, raw_text string) (go_bot.Expense, error)
	Delete(id int) error
	GetByTime(timeUnix int64) ([]go_bot.ExpenseWCN, error)
	GetByTimeByGroup(timeUnix int64) ([]go_bot.ExpenseWCN, error)
}

type Category interface {
	GetOne(string) (go_bot.Category, error)
	GetAll() ([]go_bot.Category, error)
}

type Aliases interface {
	GetOne(text string) (go_bot.Alias, error)
	GetAllInGroup(codename string) ([]go_bot.Alias, error)
	GetAllByGroups() ([]go_bot.AliasGroup, error)
	Delete(id int) error
	Create(codecame string, text string) error
}

type Repository struct {
	Aliases
	Expenses
	Category
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Aliases:  NewAliasesSqlite3(db),
		Expenses: NewExpensesSqlite3(db),
		Category: NewCategorySqlite3(db),
	}
}
