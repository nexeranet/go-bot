package service

import (
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Command interface {
	AddHandler(command string, callback interface{})
	SetDefaultHandler(callback interface{})
	Notify(update *tgbotapi.Update)
}

type Service struct {
	Command
	handlers       []*Handler
	bot            *tgbotapi.BotAPI
	DefaultHandler interface{}
}

func NewService(bot *tgbotapi.BotAPI) *Service {
	return &Service{
		bot:            bot,
		handlers:       []*Handler{},
		DefaultHandler: func(update *tgbotapi.Update) {},
	}
}

func (s *Service) AddHandler(command string, callback interface{}) {
	handler := &Handler{
		command:   command,
		callback:  callback,
		isCommand: false,
	}
	handler.Setup()
	s.handlers = append(s.handlers, handler)

}
func (s *Service) SetDefaultHandler(callback interface{}) {
	s.DefaultHandler = callback
}

func (s *Service) Notify(update *tgbotapi.Update) error {
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
