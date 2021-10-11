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

func (c *CategorySqlite3) GetOne(name string, tg_id int64) (go_bot.Category, error) {
	var category go_bot.Category
	err := c.db.QueryRow("SELECT codename, name FROM category WHERE codename=$1", name).Scan(&category.Codename, &category.Name, &category.IsBaseExpense)
	if err != nil {
		return category, err
	}
	return category, nil
}
func (c *CategorySqlite3) GetAll(tg_id int64) ([]go_bot.Category, error) {
	var categories []go_bot.Category
	query := "SELECT codename, name FROM category WHERE tg_id=$1"
	rows, err := c.db.Query(query, tg_id)
	if err != nil {
		return categories, err
	}
	defer rows.Close()
	for rows.Next() {
		var ctg go_bot.Category
		err = rows.Scan(&ctg.Codename, &ctg.Name)
		if err != nil {
			return categories, err
		}
		categories = append(categories, ctg)
	}
	return categories, nil
}

func (c *CategorySqlite3) Create(codename string, name string, tg_id int64) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	query := "INSERT INTO category (codename, name, tg_id) values ($1, $2, $3)"
	statement, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = statement.Exec(codename, name, tg_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	insert_alias_query := "INSERT INTO alias (category_codename, text, tg_id) values ($1, $2, $3)"
	statement, err = tx.Prepare(insert_alias_query)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = statement.Exec(codename, name, tg_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *CategorySqlite3) Delete(codename string, tg_id int64) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE expense SET category_codename=$1 WHERE category_codename=$2 AND tg_id=$3"
	statement, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = statement.Exec("other", codename, tg_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	del_alias_query := "DELETE FROM alias WHERE category_codename=$1 AND tg_id=$2"
	statement, err = tx.Prepare(del_alias_query)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = statement.Exec(codename, tg_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	del_category_query := "DELETE FROM category WHERE codename=$1 AND tg_id=$2"
	statement, err = tx.Prepare(del_category_query)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = statement.Exec(codename, tg_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
