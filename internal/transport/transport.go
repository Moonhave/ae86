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

func Start(config Config, restControllers *container.RestContainer, botHandlers *container.HandlerContainer) error {
	err := bot.Start(config.Bot, botHandlers)
	if err != nil {
		return err
	}

	err = rest.Start(config.Rest, restControllers)
	if err != nil {
		return err
	}

	return nil
}
