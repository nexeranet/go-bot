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
func (c *CategorySqlite3) GetAll() ([]go_bot.Category, error) {
	var categories []go_bot.Category
	query := "SELECT codename, name, is_base_expense FROM category"
	rows, err := c.db.Query(query)
	if err != nil {
		return categories, err
	}
	defer rows.Close()
	for rows.Next() {
		var ctg go_bot.Category
		err = rows.Scan(&ctg.Codename, &ctg.Name, &ctg.IsBaseExpense)
		if err != nil {
			return categories, err
		}
		categories = append(categories, ctg)
	}
	return categories, nil
}
func (c *CategorySqlite3) Delete(codename string) error {
	tx, _ := c.db.Begin()
	query := "UPDATE expense SET category_codename=$1 WHERE category_codename=$2"
	_, err := tx.Query(query, "other", codename)
	if err != nil {
		tx.Rollback()
		return err
	}
	del_alias_query := "DELETE FROM alias WHERE category_codename=$1"
	_, err = tx.Query(del_alias_query, codename)
	if err != nil {
		tx.Rollback()
		return err
	}
	del_category_query := "DELETE FROM category WHERE codename=$1"
	_, err = tx.Query(del_category_query, codename)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
