package repository

import (
	"database/sql"

	go_bot "github.com/nexeranet/go-bot"
)

type User interface {
	Create(id int64) error
	GetOne(id int64) (go_bot.TgUser, error)
}

type Expenses interface {
	Create(category string, amount int, raw_text string, tg_id int64) (go_bot.Expense, error)
	Delete(id int, tg_id int64) error
	GetByTime(timeUnix int64, tg_id int64) ([]go_bot.ExpenseWCN, error)
	GetByTimeByGroup(timeUnix int64, tg_id int64) ([]go_bot.ExpenseWCN, error)
}

type Category interface {
	GetOne(name string, tg_id int64) (go_bot.Category, error)
	GetAll(tg_id int64) ([]go_bot.Category, error)
	Delete(codename string, tg_id int64) error
	Create(codename string, name string, tg_id int64) error
}

type Aliases interface {
	GetOne(text string, tg_id int64) (go_bot.Alias, error)
	GetAllInGroup(codename string, tg_id int64) ([]go_bot.Alias, error)
	GetAllByGroups(tg_id int64) ([]go_bot.AliasGroup, error)
	Delete(id int, tg_id int64) error
	Create(codecame string, text string, tg_id int64) error
}

type Repository struct {
	User
	Aliases
	Expenses
	Category
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:     NewUsersSqlite3(db),
		Aliases:  NewAliasesSqlite3(db),
		Expenses: NewExpensesSqlite3(db),
		Category: NewCategorySqlite3(db),
	}
}
