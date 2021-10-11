package handler

import (
	"errors"
	"fmt"

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

func (h *Handler) StartCommand(update *tgbotapi.Update) error {
	err := h.CheckUser(update)
	if err != nil {
		err = h.repos.User.Create(update.Message.Chat.ID)
		if err != nil {
			h.bot.Send(err.Error(), update)
			return err
		}
	}
	str := "Бот для учёта финансов\n\nДобавить расход: 250 такси\nСегодняшняя статистика: /today\nЗа текущий месяц: /month\nПоследние внесённые расходы: /expenses\nКатегории трат: /categories"
	h.bot.Send(str, update)
	return nil
}

func (h *Handler) CheckUser(update *tgbotapi.Update) error {
	user, err := h.repos.User.GetOne(update.Message.Chat.ID)
	if err != nil {
		return err
	}
	if user.Tg_Id == 0 {
		return errors.New("Not register user")
	}
	return nil
}

func (h *Handler) AuthGuard(update *tgbotapi.Update) error {
	err := h.CheckUser(update)
	if err != nil {
		h.bot.Send("Please register /register", update)
		return err
	}
	return nil
}

func (h *Handler) RegisterUser(update *tgbotapi.Update) error {
	_, err := h.repos.User.GetOne(update.Message.Chat.ID)
	if err != nil {
		err := h.repos.User.Create(update.Message.Chat.ID)
		if err != nil {
			h.bot.Send(err.Error(), update)
			return err
		}
	}
	h.bot.Send(fmt.Sprintf("You are registered: %d", update.Message.Chat.ID), update)
	return nil
}

func (h *Handler) InitBotHandlers() {
	h.bot.AddHandler("/register", h.RegisterUser)
	h.bot.AddHandler("/start", h.StartCommand)
	h.bot.AddHandler("/today", h.AuthGuard, h.GetTodayStatistics)
	h.bot.AddHandler("/month", h.AuthGuard, h.GetMonthStatistics)
	h.bot.AddHandler("/expenses", h.AuthGuard, h.GetExpensesStatistics)
	h.bot.AddHandler("/category", h.AuthGuard, h.GetCategoryByName)
	h.bot.AddHandler("/del_category", h.AuthGuard, h.DeleteCategory)
	h.bot.AddHandler("/add_category", h.AuthGuard, h.CreateCategory)
	h.bot.AddHandler("/categories", h.AuthGuard, h.GetCategories)
	h.bot.AddHandler("/del", h.AuthGuard, h.DeleteExpense)
	h.bot.AddHandler("/del_alias", h.AuthGuard, h.DeleteAlias)
	h.bot.AddHandler("/add_alias", h.AuthGuard, h.CreateAlias)
	h.bot.SetDefaultHandler(func(update *tgbotapi.Update) {
		if update.Message.IsCommand() {
			return
		}
		err := h.AuthGuard(update)
		if err != nil {
			return
		}
		h.CreateExpense(update)
	})
}
