package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nexeranet/go-bot/pkg/repository"
)

func InitBotHandlers(bot *Bot, repos *repository.Repository) {

	bot.AddHandler("hello", func(update *tgbotapi.Update) {
		bot.Send("Hello", update)
	})
	bot.AddHandler("/start", func(update *tgbotapi.Update) {
		bot.Send("Start you work", update)
	})
	bot.AddHandler("/category", func(update *tgbotapi.Update) {
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
	})
	bot.SetDefaultHandler(func(update *tgbotapi.Update) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.bot.Send(msg); err != nil {
			panic(err)
		}
	})
}
