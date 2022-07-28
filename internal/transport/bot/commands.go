package bot

import (
	"ae86/internal/container"
	"gopkg.in/telebot.v3"
)

func RegisterCommands(bot *telebot.Bot, handlers *container.HandlerContainer) {
	// example
	bot.Handle("/categories", handlers.Category().LoadCategories)
}
