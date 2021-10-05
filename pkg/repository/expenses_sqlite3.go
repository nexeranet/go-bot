package repository

import (
	"database/sql"
)

type ExpensesSqlite3 struct {
	db *sql.DB
}

func NewExpensesSqlite3(db *sql.DB) *ExpensesSqlite3 {
	return &ExpensesSqlite3{
		db: db,
	}
}
