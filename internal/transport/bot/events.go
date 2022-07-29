package bot

import (
	"ae86/internal/container"
	tele "gopkg.in/telebot.v3"
)

func RegisterEvents(bot *tele.Bot, handlers *container.HandlerContainer) {
	bot.Handle(tele.OnText, handlers.Customer().TryStoreAddress)
}
