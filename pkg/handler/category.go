package handler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nexeranet/go-bot/pkg/bot"
	"github.com/nexeranet/go-bot/pkg/repository"
)

func GetCategoryByName(bot *bot.Bot, repos *repository.Repository, update *tgbotapi.Update) {
	argString := update.Message.CommandArguments()
	if argString == "" {
		bot.Send("No arguments", update)
		return
	}
	fmt.Printf("argString: %v\n", argString)
	category, err := repos.Category.GetOne(argString)
	if err != nil {
		bot.Send("No category with this name", update)
		return
	}
	bot.Send(fmt.Sprintf("category:\n%v", category), update)
}
func GetCategories(bot *bot.Bot, repos *repository.Repository, update *tgbotapi.Update) {
	categories, err := repos.Category.GetAll()
	if err != nil {
		bot.Send(err.Error(), update)
		return
	}
	msg := "Категории трат:\n\n"
	for _, category := range categories {
		template := fmt.Sprintf("Идентификатор %s, название: %s\n", category.Codename, category.Name)
		msg = fmt.Sprintf("%s%s", msg, template)
	}
	bot.Send(msg, update)
}
