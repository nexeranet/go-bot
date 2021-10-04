package service

import (
	"reflect"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type HandlerIntr interface {
	Validate(update *tgbotapi.Update) bool
	Notify(update *tgbotapi.Update) interface{}
	Setup() bool
}

type Handler struct {
	HandlerIntr
	command   string
	isCommand bool
	callback  interface{}
}

func (h *Handler) Validate(update *tgbotapi.Update) bool {
	if h.isCommand {
		return update.Message.Command() == h.command
	}
	return strings.Contains(update.Message.Text, h.command)
}
func (h *Handler) Notify(update *tgbotapi.Update) interface{} {
	v := reflect.ValueOf(h.callback)
	vargs := make([]reflect.Value, 1)
	vargs[0] = reflect.ValueOf(update)
	return v.Call(vargs)
}
func (h *Handler) Setup() {
	h.isCommand = strings.HasPrefix(h.command, "/")
	if h.isCommand {
		h.command = strings.Replace(h.command, "/", "", -1)
	}
}
