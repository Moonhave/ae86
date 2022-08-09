package bot

import (
	"ae86/internal/container"
	"ae86/internal/transport/bot/view"
	tele "gopkg.in/telebot.v3"
)

func RegisterButtonCallbacks(bot *tele.Bot, handlers *container.HandlerContainer) {
	bot.Handle(&view.BtnCategories, handlers.Category().SendCategories)

	bot.Handle(&view.BtnCart, handlers.Order().SendCart)
	bot.Handle(&view.BtnClearCart, handlers.Order().ClearCart)
	bot.Handle(&view.BtnOrder, handlers.Order().PromptAddressInput)
	bot.Handle(&view.BtnInlineOrder, handlers.Order().PromptAddressInput)
	bot.Handle(&view.BtnCancelOrder, handlers.Order().CancelOrder)

	bot.Handle(&view.BtnCash, handlers.Order().SetCashAsPaymentMethod)
	bot.Handle(&view.BtnCard, handlers.Order().SetCardAsPaymentMethod)

	bot.Handle(&view.BtnConfirmOrder, handlers.Order().SendOrder)
	bot.Handle(&view.BtnCancelConfirmOrder, handlers.Order().CancelOrder)

	bot.Handle(&view.BtnOrderList, handlers.Order().SendOrderList)

	bot.Handle(&view.BtnBack, handlers.General().GoBackToMenu)
	bot.Handle(&view.BtnInlineBack, handlers.General().GoBackToMenu)
}
