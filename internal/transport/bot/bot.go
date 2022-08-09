package bot

import (
	"ae86/internal/container"
	"ae86/internal/transport/bot/view"
	tele "gopkg.in/telebot.v3"
)

func Start(config Config, handlers *container.HandlerContainer) error {
	bot, err := tele.NewBot(tele.Settings{
		Token: config.Token,
		Poller: &tele.LongPoller{
			Timeout: config.LongPollerTimeout,
		},
	})
	if err != nil {
		return err
	}

	InitializeMenuReplies()
	RegisterCommands(bot, handlers)
	RegisterButtonCallbacks(bot, handlers)
	RegisterEvents(bot)

	bot.Start()
	return nil
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

	view.OrderInfoMenu.Reply(
		view.OrderInfoMenu.Row(view.BtnCancelOrder),
	)

	view.PaymentMethodMenu.Reply(
		view.PaymentMethodMenu.Row(view.BtnCard, view.BtnCash),
		view.PaymentMethodMenu.Row(view.BtnCancelOrder),
	)

	view.ConfirmOrderMenu.Reply(
		view.ConfirmOrderMenu.Row(view.BtnConfirmOrder),
		view.ConfirmOrderMenu.Row(view.BtnCancelConfirmOrder),
	)

	view.EmptyMenu.Reply(
		view.EmptyMenu.Row(view.BtnBack),
	)

	view.ProductMenu.Reply()
}
