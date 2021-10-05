package handler

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nexeranet/go-bot/pkg/bot"
	"github.com/nexeranet/go-bot/pkg/repository"
)

func GetTodayStatisticse(bot *bot.Bot, repos *repository.Repository, update *tgbotapi.Update) {
	now := time.Now()
	yesteday := now.Add(-24 * time.Hour)
	expeneses, err := repos.Expenses.GetByTime(yesteday.Unix())
	if err != nil {
		bot.Send(err.Error(), update)
	}
	sum := 0
	msg := "Расходы сегодня:\n\n"
	for _, exp := range expeneses {
		template := fmt.Sprintf("Category:%s, amount: %d \n", exp.CategoryCodename, exp.Amount)
		msg = fmt.Sprintf("%s%s", msg, template)
		sum = sum + exp.Amount
	}
	msg = fmt.Sprintf("%s\nSUM:%d\n", msg, sum)
	bot.Send(msg, update)
}
