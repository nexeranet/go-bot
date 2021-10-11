package repository

import (
	"database/sql"

	go_bot "github.com/nexeranet/go-bot"
)

type UsersSqlite3 struct {
	db *sql.DB
}

func NewUsersSqlite3(db *sql.DB) *UsersSqlite3 {
	return &UsersSqlite3{
		db: db,
	}
}

func (u *UsersSqlite3) Create(id int64) error {
	query := "INSERT INTO users(tg_id) values ($1)"
	statement, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersSqlite3) GetOne(id int64) (go_bot.TgUser, error) {
	var user go_bot.TgUser
	query := "SELECT id, tg_id FROM users WHERE tg_id=$1"
	err := u.db.QueryRow(query, id).Scan(&user.Id, &user.Tg_Id)
	if err != nil {
		return user, err
	}
	return user, nil
}
