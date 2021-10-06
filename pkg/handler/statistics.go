package handler

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func PrintStatistics(startDay int64, title string, h *Handler, update *tgbotapi.Update) {
	expeneses, err := h.repos.Expenses.GetByTime(startDay)
	if err != nil {
		h.bot.Send(err.Error(), update)
	}
	sum := 0
	msg := title + "\n\n"
	for _, exp := range expeneses {
		template := fmt.Sprintf("Категория: %s, сумма: %d (для удаления /del %d)\n", exp.CategoryName, exp.Amount, exp.Id)
		msg = fmt.Sprintf("%s%s", msg, template)
		sum = sum + exp.Amount
	}
	msg = fmt.Sprintf("%s\nOбщая сумма: %d\n", msg, sum)
	h.bot.Send(msg, update)
}
func (h *Handler) GetTodayStatistics(update *tgbotapi.Update) {
	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	PrintStatistics(startDay.Unix(), "Расходы сегодня:", h, update)
}

func (h *Handler) GetMonthStatistics(update *tgbotapi.Update) {
	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	PrintStatistics(startDay.Unix(), "Расходы за месяц:", h, update)
}
