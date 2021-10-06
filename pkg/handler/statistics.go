package handler

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	go_bot "github.com/nexeranet/go-bot"
)

func PrintStatistics(expeneses []go_bot.ExpenseWCN, title string, h *Handler, update *tgbotapi.Update) {
	sum := 0
	msg := title + "\n\n"
	for _, exp := range expeneses {
		template := fmt.Sprintf("Категория: %s, сумма: %d \n", exp.CategoryName, exp.Amount)
		msg = fmt.Sprintf("%s%s", msg, template)
		sum = sum + exp.Amount
	}
	msg = fmt.Sprintf("%s\nOбщая сумма: %d\n", msg, sum)
	h.bot.Send(msg, update)
}

func (h *Handler) GetExpensesStatistics(update *tgbotapi.Update) {
	now := time.Now()
	startDay := now.Add(-24 * time.Hour)
	expeneses, err := h.repos.Expenses.GetByTime(startDay.Unix())
	if err != nil {
		h.bot.Send(err.Error(), update)
	}
	sum := 0
	msg := "Расходы за 24 часа:" + "\n\n"
	for _, exp := range expeneses {
		template := fmt.Sprintf("Категория: %s, сумма: %d \n", exp.CategoryName, exp.Amount)
		msg = fmt.Sprintf("%s%s", msg, template)
		sum = sum + exp.Amount
	}
	msg = fmt.Sprintf("%s\nOбщая сумма: %d\n", msg, sum)
	h.bot.Send(msg, update)
	//PrintStatistics(expeneses, "Расходы за 24 часа:", h, update)
}

func (h *Handler) GetTodayStatistics(update *tgbotapi.Update) {
	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	expeneses, err := h.repos.Expenses.GetByTimeByGroup(startDay.Unix())
	if err != nil {
		h.bot.Send(err.Error(), update)
	}
	PrintStatistics(expeneses, "Расходы сегодня:", h, update)
}

func (h *Handler) GetMonthStatistics(update *tgbotapi.Update) {
	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	expeneses, err := h.repos.Expenses.GetByTimeByGroup(startDay.Unix())
	if err != nil {
		h.bot.Send(err.Error(), update)
	}
	PrintStatistics(expeneses, "Расходы за месяц:", h, update)
}
