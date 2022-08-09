package bot

import (
	"ae86/internal/container"
	"ae86/internal/transport/bot/view"
	"ae86/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

func Start(config Config, handlers *container.HandlerContainer) {
	bot, err := tele.NewBot(tele.Settings{
		Token: config.Token,
		Poller: &tele.LongPoller{
			Timeout: config.LongPollerTimeout,
		},
	})
	if err != nil {
		logger.Log.Fatalf("failed to create bot: %v", err)
	}

	InitializeMenuReplies()
	RegisterCommands(bot, handlers)
	RegisterButtonCallbacks(bot, handlers)
	RegisterEvents(bot)

	bot.Start()
}

func InitializeMenuReplies() {
	view.Menu.Reply(
		view.Menu.Row(view.BtnCategories),
		view.Menu.Row(view.BtnCart, view.BtnOrder),
		view.Menu.Row(view.BtnInfo),
		view.Menu.Row(view.BtnOrderList),
		view.Menu.Row(view.BtnContactManager),
	)

	view.CartMenu.Reply(
		view.CartMenu.Row(view.BtnClearCart),
		view.CartMenu.Row(view.BtnCategoryBack, view.BtnOrder),
	)

	view.AddressMenu.Reply(
		view.AddressMenu.Row(view.BtnCancelOrder),
	)

	view.PaymentMethodMenu.Reply(
		view.PaymentMethodMenu.Row(view.BtnCard, view.BtnCash),
		view.PaymentMethodMenu.Row(view.BtnCancelOrder),
	)

	view.EmptyMenu.Reply(
		view.EmptyMenu.Row(view.BtnBack),
	)

	view.ProductMenu.Reply()
}
