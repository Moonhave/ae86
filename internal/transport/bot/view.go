package bot

import tele "gopkg.in/telebot.v3"

var (
	menu              = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnCategories     = menu.Text("Меню")
	btnCart           = menu.Text("Корзина")
	btnOrder          = menu.Text("Оформить заказ")
	btnInfo           = menu.Text("О нас")
	btnOrderList      = menu.Text("Мои заказы")
	btnContactManager = menu.Text("Связаться с менеджером")

	menuMessage         = "Главное меню"
	infoMessage         = "Здесь пока что пусто)"
	managerMessage      = "Контакт менеджера: @danqzq"
	emptyMessage        = "Пусто"
	orderMessage        = "Заказ оформлен"
	selectAmountMessage = "Выберите количество"
	addedToCartMessage  = "Товар добавлен в корзину"
	cartEmptyMessage    = "Корзина пуста"

	categoryMenuRows    []tele.Row
	categoryMenu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	categoryMenuMessage = "Выберите категорию:"
	btnCategoryBack     = categoryMenu.Text("Назад в главное меню")

	cartMenu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnClearCart = cartMenu.Text("Очистить корзину")

	productMenu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnInlineAddMessage = "Добавить"
	btnInlineAdded      = productMenu.Data("Добавлено в корзину", "added")
	btnInlineOrder      = productMenu.Data("Оформить заказ", "order")
	btnInlineBack       = productMenu.Data("Вернуться в главное меню", "back")

	addressMenu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	addressMenuMessage = "Введите адрес доставки:"
	btnCancelOrder     = addressMenu.Text("Отмена")

	paymentMethodMenu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	paymentMethodMenuMessage = "Выберите способ оплаты:"
	btnCreditCard            = paymentMethodMenu.Text("Кредитная карта")
	btnCash                  = paymentMethodMenu.Text("Наличными")

	emptyMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnBack   = emptyMenu.Text("Назад в главное меню")

	unknownCommandMessage = "Неизвестная команда"
)
