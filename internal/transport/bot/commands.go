package bot

import (
	"ae86/internal/container"
	tele "gopkg.in/telebot.v3"
)

func RegisterCommands(bot *tele.Bot, handlers *container.HandlerContainer) {
	bot.Handle("/start", handlers.General().Start)

	bot.Handle("/categories", handlers.Category().SendCategories)
}
