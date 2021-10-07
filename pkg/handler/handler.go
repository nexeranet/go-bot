package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nexeranet/go-bot/pkg/bot"
	"github.com/nexeranet/go-bot/pkg/repository"
)

type Handler struct {
	bot   *bot.Bot
	repos *repository.Repository
}

func NewHandler(bot *bot.Bot, repos *repository.Repository) *Handler {
	return &Handler{
		bot:   bot,
		repos: repos,
	}
}

func (h *Handler) StartCommand(update *tgbotapi.Update) {
	str := "Бот для учёта финансов\n\nДобавить расход: 250 такси\nСегодняшняя статистика: /today\nЗа текущий месяц: /month\nПоследние внесённые расходы: /expenses\nКатегории трат: /categories"
	h.bot.Send(str, update)
}

func (h *Handler) InitBotHandlers() {
	h.bot.AddHandler("/start", h.StartCommand)
	h.bot.AddHandler("/today", h.GetTodayStatistics)
	h.bot.AddHandler("/month", h.GetMonthStatistics)
	h.bot.AddHandler("/expenses", h.GetExpensesStatistics)
	h.bot.AddHandler("/category", h.GetCategoryByName)
	h.bot.AddHandler("/del_category", h.DeleteCategory)
	h.bot.AddHandler("/add_category", h.CreateCategory)
	h.bot.AddHandler("/categories", h.GetCategories)
	h.bot.AddHandler("/del", h.DeleteExpense)
	h.bot.AddHandler("/del_alias", h.DeleteAlias)
	h.bot.AddHandler("/add_alias", h.CreateAlias)
	h.bot.SetDefaultHandler(func(update *tgbotapi.Update) {
		if update.Message.IsCommand() {
			return
		}
		h.CreateExpense(update)
	})
}
