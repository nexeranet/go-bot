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

func (e *ExpensesSqlite3) Create(category string, amount int, raw_text string, tg_id int64) (go_bot.Expense, error) {
	var expense go_bot.Expense
	query := "INSERT INTO expense (amount, created, category_codename, raw_text, tg_id) values($1, $2, $3, $4, $5)"
	statement, err := e.db.Prepare(query)
	if err != nil {
		return expense, err
	}
	_, err = statement.Exec(amount, time.Now().Unix(), category, raw_text, tg_id)
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func (e *ExpensesSqlite3) Delete(id int, tg_id int64) error {
	query := "DELETE FROM expense WHERE id=$1 AND tg_id=$2"
	statement, err := e.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id, tg_id)
	if err != nil {
		return err
	}
	return nil
}

func (e *ExpensesSqlite3) GetByTime(timeUnix int64, tg_id int64) ([]go_bot.ExpenseWCN, error) {
	var expenses []go_bot.ExpenseWCN
	query := "SELECT id, amount, created, category_codename, name FROM expense INNER JOIN category on expense.category_codename = category.codename WHERE created >= $1 AND expense.tg_id=$2"
	rows, err := e.db.Query(query, timeUnix, tg_id)
	if err != nil {
		return expenses, err
	}

	defer rows.Close()
	for rows.Next() {
		var exp go_bot.ExpenseWCN
		err = rows.Scan(&exp.Id, &exp.Amount, &exp.Created, &exp.CategoryCodename, &exp.CategoryName)
		if err != nil {
			return expenses, err
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}
func (e *ExpensesSqlite3) GetByTimeByGroup(timeUnix int64, tg_id int64) ([]go_bot.ExpenseWCN, error) {
	var expenses []go_bot.ExpenseWCN
	query := "SELECT id, SUM(amount) as sum, created, category_codename, name FROM expense INNER jOIN category on category.codename= expense.category_codename WHERE created >= $1 AND expense.tg_id=$2 GROUP BY category_codename"
	rows, err := e.db.Query(query, timeUnix, tg_id)
	if err != nil {
		return expenses, err
	}

	defer rows.Close()
	for rows.Next() {
		var exp go_bot.ExpenseWCN
		err = rows.Scan(&exp.Id, &exp.Amount, &exp.Created, &exp.CategoryCodename, &exp.CategoryName)
		if err != nil {
			return expenses, err
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}
