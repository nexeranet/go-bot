package main

import (
	"database/sql"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nexeranet/go-bot/pkg/bot"
	"github.com/nexeranet/go-bot/pkg/handler"
	"github.com/nexeranet/go-bot/pkg/repository"
)

// @TODO refactor my custom bot, make dependancy injection
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database, _ := sql.Open("sqlite3", "./db/db.db")
	repos := repository.NewRepository(database)
	tgbot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	tgbot.Debug = true

	log.Printf("Authorized on account %s", tgbot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tgbot.GetUpdatesChan(u)
	if err != nil {
		log.Panic("Error on Get Updates Chan")
	}
	cBot := bot.NewBot(tgbot)
	handl := handler.NewHandler(cBot, repos)
	handl.InitBotHandlers()

	// userId := os.Getenv("TELEGRAM_ID")
	// id, err := strconv.ParseInt(userId, 10, 64)
	// if err != nil {
	// log.Panic("Error env TELEGRAM_ID empty or not number")
	// }
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		// if update.Message.Chat.ID != id {
		// continue
		// }
		err := cBot.Notify(&update)
		if err != nil {
			log.Panic("Notify failed")
		}
	}
}
