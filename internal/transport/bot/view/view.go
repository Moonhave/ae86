package view

import tele "gopkg.in/telebot.v3"

var (
	Menu              = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnCategories     = Menu.Text("📋 Меню")
	BtnCart           = Menu.Text("🛒 Корзина")
	BtnOrder          = Menu.Text("🚘 Оформить заказ")
	BtnInfo           = Menu.Text("💬 О нас")
	BtnOrderList      = Menu.Text("📝 Мои заказы")
	BtnContactManager = Menu.Text("🧑‍💻 Связаться с менеджером")

	MenuMessage           = "Главное меню"
	InfoMessage           = "Здесь пока что пусто)"
	DefaultManagerMessage = "Контакт менеджера: @danqzq"
	EmptyMessage          = "Пусто"
	ConfirmOrderMessage   = "Вы уверены, что хотите оформить заказ?"
	OrderMessage          = "Заказ оформлен"
	SelectAmountMessage   = "Выберите количество"
	AddedToCartMessage    = "Товар добавлен в корзину"
	CartEmptyMessage      = "Корзина пуста"

	CategoryMenu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	CategoryMenuMessage = "Выберите категорию:"
	BtnCategoryBack     = CategoryMenu.Text("📋 Назад в главное меню")

	CartMenu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnClearCart = CartMenu.Text("❌ Очистить корзину")

	ProductMenu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnInlineAddMessage = "🔘 Добавить"
	BtnInlineAdded      = ProductMenu.Data("✅ Добавлено в корзину", "added")
	BtnInlineOrder      = ProductMenu.Data("🚘 Оформить заказ", "order")
	BtnInlineBack       = ProductMenu.Data("📋 Вернуться в главное меню", "back")

	OrderInfoMenu      = &tele.ReplyMarkup{ResizeKeyboard: true}
	AddressMenuMessage = "Введите адрес доставки:"
	BtnCancelOrder     = OrderInfoMenu.Text("❌ Отмена")

	PhoneMenuMessage = "Введите ваш номер телефона:"

	PaymentMethodMenu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	PaymentMethodMenuMessage = "Выберите способ оплаты:"
	BtnCard                  = PaymentMethodMenu.Text("💳 Кредитная карта")
	BtnCash                  = PaymentMethodMenu.Text("💵 Наличными")

	ConfirmOrderMenu      = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnConfirmOrder       = ConfirmOrderMenu.Text("✅ Оформить заказ")
	BtnCancelConfirmOrder = ConfirmOrderMenu.Text("❌ Отмена")

	EmptyMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnBack   = EmptyMenu.Text("📋 Назад в главное меню")

	UnknownCommandMessage = "❗ Неизвестная команда"
)
