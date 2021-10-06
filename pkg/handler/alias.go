package handler

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) CreateAlias(update *tgbotapi.Update) {
	argString := update.Message.CommandArguments()
	if argString == "" {
		h.bot.Send("No arguments", update)
		return
	}
	group := getParams(`(?P<Codename>[\w ]+) (?P<Text>[\w ]*)`, argString)
	if group["Codename"] == "" || group["Text"] == "" {
		h.bot.Send("Invalid arguments", update)
		return
	}
	_, err := h.repos.Category.GetOne(group["Codename"])
	if err != nil {
		h.bot.Send(err.Error(), update)
		return
	}
	err = h.repos.Aliases.Create(group["Codename"], group["Text"])
	if err != nil {
		h.bot.Send(err.Error(), update)
		return
	}
	h.bot.Send("Алеас создан", update)
}
func (h *Handler) DeleteAlias(update *tgbotapi.Update) {
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
	err = h.repos.Aliases.Delete(id)
	if err != nil {
		h.bot.Send(err.Error(), update)
		return
	}
	h.bot.Send("Удалил", update)

}
