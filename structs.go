package go_bot

type Expense struct {
	Id               int    `db:"id"`
	Amount           int    `db:"amount"`
	Created          string `db:"created"`
	CategoryCodename string `db:"category_codename"`
	RawText          string `db:"raw_text"`
}
type Category struct {
	Codename      string `db:"codename"`
	Name          string `db:"name"`
	IsBaseExpense bool   `db:"is_base_expense"`
}
