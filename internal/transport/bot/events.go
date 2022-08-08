package bot

import (
	"ae86/internal/transport/bot/view"
	tele "gopkg.in/telebot.v3"
)

func RegisterEvents(bot *tele.Bot) {
	bot.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(view.UnknownCommandMessage, view.EmptyMenu)
	})
}
