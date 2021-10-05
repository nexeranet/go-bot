package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nexeranet/go-bot/pkg/bot"
	"github.com/nexeranet/go-bot/pkg/repository"
)

func InitBotHandlers(bot *bot.Bot, repos *repository.Repository) {
	bot.AddHandler("hello", func(update *tgbotapi.Update) {
		bot.Send("Hello", update)
	})
	bot.AddHandler("/start", func(update *tgbotapi.Update) {
		str := "Бот для учёта финансов\n\nДобавить расход: 250 такси\nСегодняшняя статистика: /today\nЗа текущий месяц: /month\nПоследние внесённые расходы: /expenses\nКатегории трат: /categories"
		bot.Send(str, update)
	})
	bot.AddHandler("/category", func(update *tgbotapi.Update) {
		GetCategoryByName(bot, repos, update)
	})
	bot.AddHandler("/categories", func(update *tgbotapi.Update) {
		GetCategories(bot, repos, update)
	})
	bot.AddHandler("/del", func(update *tgbotapi.Update) {
		DeleteExpense(bot, repos, update)
	})
	bot.SetDefaultHandler(func(update *tgbotapi.Update) {
		if update.Message.IsCommand() {
			return
		}
		CreateExpense(bot, repos, update)
	})
}
