package transport

import (
	"ae86/internal/container"
	"ae86/internal/transport/bot"
	"ae86/internal/transport/rest"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func Start(conf rest.Config, restContainer *container.RestContainer) error {
	// telegram bot start
	pref := tele.Settings{
		Token:  conf.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bot.LoadCategories(b)
	bot.InitializeMenuReplies()
	bot.RegisterEndpointCallbacks(b)
	bot.RegisterButtonCallbacks(b)

	b.Start()

	err = rest.Start(conf, restContainer)
	if err != nil {
		return err
	}

	return nil
}
