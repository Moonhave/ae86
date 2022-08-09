package transport

import (
	"ae86/internal/container"
	"ae86/internal/transport/bot"
	"ae86/internal/transport/rest"
)

type Config struct {
	Rest rest.Config
	Bot  bot.Config
}

func Start(config Config, restControllers *container.RestContainer, botHandlers *container.HandlerContainer) {
	go bot.Start(config.Bot, botHandlers)
	go rest.Start(config.Rest, restControllers)
}
