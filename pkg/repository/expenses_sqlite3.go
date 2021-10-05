package repository

import (
	"database/sql"
	"time"

	go_bot "github.com/nexeranet/go-bot"
)

type ExpensesSqlite3 struct {
	db *sql.DB
}

func NewExpensesSqlite3(db *sql.DB) *ExpensesSqlite3 {
	return &ExpensesSqlite3{
		db: db,
	}
}

func (e *ExpensesSqlite3) Create(category string, amount int, raw_text string) (go_bot.Expense, error) {
	var expense go_bot.Expense
	query := "INSERT INTO expense (amount, created, category_codename, raw_text) values($1, $2, $3, $4)"
	statement, err := e.db.Prepare(query)
	if err != nil {
		return expense, err
	}
	_, err = statement.Exec(amount, time.Now().Unix(), category, raw_text)
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func (e *ExpensesSqlite3) Delete(id int) error {
	query := "DELETE FROM expense WHERE id=$1"
	statement, err := e.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (e *ExpensesSqlite3) GetByTime(timeUnix int64) ([]go_bot.Expense, error) {
	var expenses []go_bot.Expense
	query := "SELECT id, amount, created, category_codename FROM expense WHERE created >= $1"
	rows, err := e.db.Query(query, timeUnix)
	if err != nil {
		return expenses, err
	}

	defer rows.Close()
	for rows.Next() {
		var exp go_bot.Expense
		err = rows.Scan(&exp.Id, &exp.Amount, &exp.Created, &exp.CategoryCodename)
		if err != nil {
			return expenses, err
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}
