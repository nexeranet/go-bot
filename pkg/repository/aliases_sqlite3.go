package repository

import (
	"database/sql"

	go_bot "github.com/nexeranet/go-bot"
)

type AliasesSqlite3 struct {
	db *sql.DB
}

func NewAliasesSqlite3(db *sql.DB) *AliasesSqlite3 {
	return &AliasesSqlite3{
		db: db,
	}
}

func (h *AliasesSqlite3) GetOne(text string, tg_id int64) (go_bot.Alias, error) {
	var alias go_bot.Alias
	query := "SELECT id, text, category_codename FROM alias WHERE text=$1 AND tg_id=$2"
	err := h.db.QueryRow(query, text, tg_id).Scan(&alias.Id, &alias.Text, &alias.CategoryCodename)
	if err != nil {
		return alias, err
	}
	return alias, nil
}

func (h *AliasesSqlite3) GetAllInGroup(codename string, tg_id int64) ([]go_bot.Alias, error) {
	var aliases []go_bot.Alias
	query := "SELECT id, text, category_codename, name FROM alias INNER JOIN category on category.codename=alias.category_codename WHERE category_codename=$1 AND alias.tg_id=$2"
	rows, err := h.db.Query(query, codename, tg_id)
	if err != nil {
		return aliases, err
	}

	defer rows.Close()
	for rows.Next() {
		var als go_bot.Alias
		err = rows.Scan(&als.Id, &als.Text, &als.CategoryCodename, &als.CategoryName)
		if err != nil {
			return aliases, err
		}
		aliases = append(aliases, als)
	}
	return aliases, nil
}

func (h *AliasesSqlite3) GetAllByGroups(tg_id int64) ([]go_bot.AliasGroup, error) {
	var aliases []go_bot.AliasGroup
	query := "SELECT codename, name, group_concat(text, ', ') AS aliases FROM category LEFT JOIN alias on category.codename = alias.category_codename WHERE category.tg_id=$1 GROUP BY category_codename"
	rows, err := h.db.Query(query, tg_id)
	if err != nil {
		return aliases, err
	}

	defer rows.Close()
	for rows.Next() {
		var als go_bot.AliasGroup
		err = rows.Scan(&als.CategoryCodename, &als.Name, &als.List)
		if err != nil {
			return aliases, err
		}
		aliases = append(aliases, als)
	}
	return aliases, nil
}

func (h *AliasesSqlite3) Delete(id int, tg_id int64) error {
	query := "DELETE FROM alias WHERE id=$1 AND tg_id=$2"
	statement, err := h.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id, tg_id)
	if err != nil {
		return err
	}
	return nil
}

func (h *AliasesSqlite3) Create(codename string, text string, tg_id int64) error {
	query := "INSERT INTO alias(category_codename, text, tg_id) values($1, $2, $3)"
	statement, err := h.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(codename, text, tg_id)
	if err != nil {
		return err
	}
	return nil
}
