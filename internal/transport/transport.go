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

func Start(config Config, restContainer *container.RestContainer) error {
	err := bot.Start(config.Bot)
	if err != nil {
		return err
	}

	err = rest.Start(config.Rest, restContainer)
	if err != nil {
		return err
	}

	return nil
}
