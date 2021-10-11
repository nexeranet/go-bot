package handler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) GetCategoryByName(update *tgbotapi.Update) {
	argString := update.Message.CommandArguments()
	if argString == "" {
		h.bot.Send("No arguments", update)
		return
	}
	fmt.Printf("argString: %v\n", argString)
	aliases, err := h.repos.Aliases.GetAllInGroup(argString, update.Message.Chat.ID)
	if err != nil {
		h.bot.Send(err.Error(), update)
		return
	}
	msg := fmt.Sprintf("Категория [%s]: \n\n", argString)
	for _, alias := range aliases {
		template := fmt.Sprintf("Алиас: %s /del_alias %d\n", alias.Text, alias.Id)
		msg = fmt.Sprintf("%s%s", msg, template)
	}
	h.bot.Send(msg, update)
}
func (h *Handler) GetCategories(update *tgbotapi.Update) {
	categories, err := h.repos.Aliases.GetAllByGroups(update.Message.Chat.ID)
	if err != nil {
		h.bot.Send(err.Error(), update)
		return
	}
	msg := "Категории трат:\n\n"
	for _, category := range categories {
		if category.List == nil {
			category.List = "Алеасов нету"
		}
		template := fmt.Sprintf("Идентификатор %s, название: %s (%s)\n", category.CategoryCodename, category.Name, category.List)
		msg = fmt.Sprintf("%s%s", msg, template)
	}
	h.bot.Send(msg, update)
}

func (h *Handler) CreateCategory(update *tgbotapi.Update) {
	argString := update.Message.CommandArguments()
	if argString == "" {
		h.bot.Send("No arguments", update)
		return
	}
	group := getParams(`(?P<Codename>[\w ]+) (?P<Name>.*)`, argString)
	if group["Codename"] == "" || group["Name"] == "" {
		h.bot.Send("Invalid arguments", update)
		return
	}
	errs := h.repos.Category.Create(group["Codename"], group["Name"], update.Message.Chat.ID)
	if errs != nil {
		h.bot.Send(errs.Error(), update)
		return
	}
	h.bot.Send("Категория создана", update)
}

func (h *Handler) DeleteCategory(update *tgbotapi.Update) {
	argString := update.Message.CommandArguments()
	if argString == "" {
		h.bot.Send("No arguments", update)
		return
	}
	group := getParams(`(?P<Codename>[\w ]*)`, argString)
	if group["Codename"] == "" || group["Codename"] == "other" {
		h.bot.Send("Invalid arguments", update)
		return
	}
	err := h.repos.Category.Delete(group["Codename"], update.Message.Chat.ID)
	if err != nil {
		h.bot.Send(err.Error(), update)
		return
	}
	h.bot.Send("Удалил", update)
}
