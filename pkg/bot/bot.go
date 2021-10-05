package bot

import (
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Command interface {
	AddHandler(command string, callback interface{})
	SetDefaultHandler(callback interface{})
	Notify(update *tgbotapi.Update)
	Send(message string, update *tgbotapi.Update)
}

type Bot struct {
	Command
	handlers       []*Handler
	bot            *tgbotapi.BotAPI
	DefaultHandler interface{}
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot:            bot,
		handlers:       []*Handler{},
		DefaultHandler: func(update *tgbotapi.Update) {},
	}
}

func (s *Bot) AddHandler(command string, callback interface{}) {
	handler := &Handler{
		command:   command,
		callback:  callback,
		isCommand: false,
	}
	handler.Setup()
	s.handlers = append(s.handlers, handler)

}
func (s *Bot) SetDefaultHandler(callback interface{}) {
	s.DefaultHandler = callback
}

func (s *Bot) Notify(update *tgbotapi.Update) error {
	for _, handler := range s.handlers {
		if handler.Validate(update) {
			handler.Notify(update)
			return nil
		}
	}
	v := reflect.ValueOf(s.DefaultHandler)
	vargs := make([]reflect.Value, 1)
	vargs[0] = reflect.ValueOf(update)
	v.Call(vargs)
	return nil
}
func (s *Bot) Send(message string, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	if _, err := s.bot.Send(msg); err != nil {
		panic(err)
	}
}
