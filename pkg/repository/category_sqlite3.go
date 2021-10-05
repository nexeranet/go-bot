package repository

import (
	"database/sql"

	go_bot "github.com/nexeranet/go-bot"
)

type CategorySqlite3 struct {
	db *sql.DB
}

func NewCategorySqlite3(db *sql.DB) *CategorySqlite3 {
	return &CategorySqlite3{
		db: db,
	}
}

func (c *CategorySqlite3) GetOne(name string) (go_bot.Category, error) {
	var category go_bot.Category
	err := c.db.QueryRow("SELECT codename, name, is_base_expense FROM category WHERE codename=$1", name).Scan(&category.Codename, &category.Name, &category.IsBaseExpense)
	if err != nil {
		return category, err
	}
	return category, nil
}
