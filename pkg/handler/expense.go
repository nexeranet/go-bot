package handler

import (
	"regexp"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/**
 * Parses url with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func getParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
func (h *Handler) CreateExpense(update *tgbotapi.Update) {
	group := getParams(`(?P<Amount>[\d ]+) (?P<Category>.*)`, update.Message.Text)
	if group["Category"] == "" || group["Amount"] == "" {
		h.bot.Send("Invalid arguments", update)
		return
	}
	amount, err := strconv.Atoi(group["Amount"])
	if err != nil {
		h.bot.Send("Invalid arguments", update)
		return
	}
	category, err := h.repos.Aliases.GetOne(group["Category"], update.Message.Chat.ID)
	if err != nil {
		category.CategoryCodename = "other"
	}
	_, errs := h.repos.Expenses.Create(category.CategoryCodename, amount, update.Message.Text, update.Message.Chat.ID)
	if errs != nil {
		h.bot.Send(errs.Error(), update)
		return
	}
	h.bot.Send("Добавлено", update)
}

func (h *Handler) DeleteExpense(update *tgbotapi.Update) {
	argString := update.Message.CommandArguments()
	if argString == "" {
		h.bot.Send("No arguments", update)
		return
	}
	group := getParams(`(?P<Id>[\d ]+)`, argString)

	id, err := strconv.Atoi(group["Id"])
	if err != nil {
		h.bot.Send("Invalid arguments", update)
		return
	}
	err = h.repos.Expenses.Delete(id, update.Message.Chat.ID)
	if err != nil {
		h.bot.Send(err.Error(), update)
		return
	}
	h.bot.Send("Удалил", update)
}
