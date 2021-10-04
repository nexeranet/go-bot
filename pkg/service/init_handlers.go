package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitServiceHandlers(srv *Service) {

	srv.AddHandler("hello", func(update *tgbotapi.Update) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello ueba")
		if _, err := srv.bot.Send(msg); err != nil {
			panic(err)
		}
	})
	srv.AddHandler("/start", func(update *tgbotapi.Update) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "start command")
		if _, err := srv.bot.Send(msg); err != nil {
			panic(err)
		}
	})
	srv.SetDefaultHandler(func(update *tgbotapi.Update) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := srv.bot.Send(msg); err != nil {
			panic(err)
		}
	})
}
