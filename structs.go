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

type Alias struct {
	Id               int    `db:"id"`
	CategoryCodename string `db:"category_codename"`
	Text             string `db:"text"`
	CategoryName     string
}
type AliasGroup struct {
	CategoryCodename string
	List             interface{}
	Name             string
}
type ExpenseWCN struct {
	Id               int    `db:"id"`
	Amount           int    `db:"amount"`
	Created          string `db:"created"`
	CategoryCodename string `db:"category_codename"`
	CategoryName     string `db:"category_name"`
}
