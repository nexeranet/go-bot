package bot

import (
	"reflect"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type HandlerIntr interface {
	Validate(update *tgbotapi.Update) bool
	Notify(update *tgbotapi.Update) interface{}
	Setup()
}

type Handler struct {
	HandlerIntr
	command   string
	isCommand bool
	callbacks []interface{}
}

func (h *Handler) Validate(update *tgbotapi.Update) bool {
	if h.isCommand {
		return update.Message.Command() == h.command
	}
	return strings.Contains(update.Message.Text, h.command)
	// valid, err := regexp.MatchString(`\s+?`+h.command+`\s+?`, update.Message.Text)
	// if err != nil {
	// panic(err.Error())
	// }
	//return valid
}
func (h *Handler) Notify(update *tgbotapi.Update) interface{} {
Loop:
	for _, callback := range h.callbacks {
		v := reflect.ValueOf(callback)
		t := reflect.TypeOf(callback)
		vargs := make([]reflect.Value, t.NumIn())
		for key := range vargs {
			vargs[key] = reflect.ValueOf(update)
		}
		values := v.Call(vargs)
		for _, val := range values {
			if !val.IsNil() {
				break Loop
			}
		}
	}
	return nil
}
func (h *Handler) Setup() {
	h.isCommand = strings.HasPrefix(h.command, "/")
	if h.isCommand {
		h.command = strings.Replace(h.command, "/", "", -1)
	}
}
