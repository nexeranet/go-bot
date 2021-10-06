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
	category, err := h.repos.Category.GetOne(argString)
	if err != nil {
		h.bot.Send("No category with this name", update)
		return
	}
	h.bot.Send(fmt.Sprintf("category:\n%v", category), update)
}
func (h *Handler) GetCategories(update *tgbotapi.Update) {
	categories, err := h.repos.Category.GetAll()
	if err != nil {
		h.bot.Send(err.Error(), update)
		return
	}
	msg := "Категории трат:\n\n"
	for _, category := range categories {
		template := fmt.Sprintf("Идентификатор %s, название: %s\n", category.Codename, category.Name)
		msg = fmt.Sprintf("%s%s", msg, template)
	}
	h.bot.Send(msg, update)
}
