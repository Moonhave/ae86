package bot

import (
	"ae86/internal/model"
	"fmt"
	tele "gopkg.in/telebot.v3"
)

var (
	menu              = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnCategories     = menu.Text("Меню")
	btnCart           = menu.Text("Корзина")
	btnOrder          = menu.Text("Оформить заказ")
	btnInfo           = menu.Text("О нас")
	btnOrderList      = menu.Text("Мои заказы")
	btnContactManager = menu.Text("Связаться с менеджером")

	menuMessage      = "Главное меню"
	infoMessage      = "Здесь пока что пусто)"
	managerMessage   = "Контакт менеджера: @danqzq"
	emptyMessage     = "Пусто"
	orderMessage     = "Заказ оформлен"
	cartEmptyMessage = "Корзина пуста"

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

type TempUserInfo struct {
	Cart             []*model.Product
	IsSettingAddress bool
}

// temp storage
var userStorage = make(map[int64]*TempUserInfo)

func getCurrentUser(c tele.Context) *TempUserInfo {
	if userStorage[c.Sender().ID] == nil {
		userStorage[c.Sender().ID] = &TempUserInfo{}
	}
	return userStorage[c.Sender().ID]
}

func LoadCategories(bot *tele.Bot) {
	// TODO: populate with actual categories from db
	var categories = []model.Category{{Title: "Пицца"}, {Title: "Суши"}, {Title: "Десерты"}, {Title: "Напитки"}}

	for _, category := range categories {
		btn := categoryMenu.Text(category.Title)
		categoryMenuRows = append(categoryMenuRows, categoryMenu.Row(btn))

		bot.Handle(&btn, func(c tele.Context) error {
			c.Send(loadProducts(bot, category), productMenu)
			return c.Respond()
		})
	}

	categoryMenuRows = append(categoryMenuRows, categoryMenu.Row(btnCategoryBack))
}

func loadProducts(bot *tele.Bot, category model.Category) string {
	// TODO: get products from db depending on category

	var testProducts = []model.Product{
		{
			Title:       category.Title + " 1",
			Description: "Описание",
			Price:       1490,
		},
		{
			Title:       category.Title + " 2",
			Description: "Описание",
			Price:       1490,
		},
	}

	var text string
	for i, product := range testProducts {
		btnAddToCart := productMenu.Data(btnInlineAddMessage, "add_product_"+fmt.Sprintf("%d", i))

		var buttonRows = []tele.Row{
			productMenu.Row(btnAddToCart),
		}
		var isLastButton = i == len(testProducts)-1
		if isLastButton {
			buttonRows = append(buttonRows, productMenu.Row(btnInlineBack))
		}
		productMenu.Inline(buttonRows...)

		text = fmt.Sprintf("%s\n%s\nЦена: %d тенге", product.Title, product.Description, product.Price)
		var p = product
		bot.Handle(&btnAddToCart, func(c tele.Context) error {
			getCurrentUser(c).Cart = append(getCurrentUser(c).Cart, &p)
			buttonRows = []tele.Row{
				productMenu.Row(btnInlineAdded),
				productMenu.Row(btnInlineOrder),
			}
			if isLastButton {
				buttonRows = append(buttonRows, productMenu.Row(btnInlineBack))
			}
			productMenu.Inline(buttonRows...)
			c.Edit(productMenu)

			return c.Respond(&tele.CallbackResponse{Text: "Товар добавлен в корзину"})
		})
	}

	return text
}

func InitializeMenuReplies() {
	menu.Reply(
		menu.Row(btnCategories),
		menu.Row(btnCart, btnOrder),
		menu.Row(btnInfo),
		menu.Row(btnOrderList),
		menu.Row(btnContactManager),
	)

	categoryMenu.Reply(categoryMenuRows...)

	cartMenu.Reply(
		cartMenu.Row(btnClearCart),
		cartMenu.Row(btnCategoryBack, btnOrder),
	)

	addressMenu.Reply(
		addressMenu.Row(btnCancelOrder),
	)

	paymentMethodMenu.Reply(
		paymentMethodMenu.Row(btnCreditCard, btnCash),
		paymentMethodMenu.Row(btnCancelOrder),
	)

	emptyMenu.Reply(
		emptyMenu.Row(btnBack),
	)

	productMenu.Reply()
}

func RegisterEndpointCallbacks(bot *tele.Bot) {
	bot.Handle("/start", func(c tele.Context) error {
		userStorage[c.Sender().ID] = &TempUserInfo{
			Cart:             []*model.Product{},
			IsSettingAddress: false,
		}
		return c.Send(menuMessage, menu)
	})
}

func RegisterButtonCallbacks(bot *tele.Bot) {
	bot.Handle(&btnCategories, func(c tele.Context) error {
		return c.Send(categoryMenuMessage, categoryMenu)
	})

	bot.Handle(&btnCart, func(c tele.Context) error {
		if len(getCurrentUser(c).Cart) == 0 {
			return c.Send(cartEmptyMessage, emptyMenu)
		}
		text := ""
		for _, product := range getCurrentUser(c).Cart {
			text += fmt.Sprintf("%s\n%s\nЦена: %d тенге\n\n", product.Title, product.Description, product.Price)
		}
		text += "Сумма: " + fmt.Sprintf("%d", priceSum(getCurrentUser(c).Cart)) + " тенге"
		return c.Send(text, cartMenu)
	})

	bot.Handle(&btnOrder, func(c tele.Context) error {
		if len(getCurrentUser(c).Cart) == 0 {
			return c.Send(cartEmptyMessage, emptyMenu)
		}

		getCurrentUser(c).IsSettingAddress = true

		return c.Send(addressMenuMessage, addressMenu)
	})

	bot.Handle(&btnInlineOrder, func(c tele.Context) error {
		if len(getCurrentUser(c).Cart) == 0 {
			return c.Send(cartEmptyMessage, emptyMenu)
		}

		getCurrentUser(c).IsSettingAddress = true

		c.Send(addressMenuMessage, addressMenu)
		return c.Respond()
	})

	bot.Handle(&btnCancelOrder, func(c tele.Context) error {
		getCurrentUser(c).IsSettingAddress = false

		return c.Send(menuMessage, menu)
	})

	bot.Handle(tele.OnText, func(c tele.Context) error {
		if !getCurrentUser(c).IsSettingAddress {
			return c.Send(unknownCommandMessage, emptyMenu)
		}

		if len(getCurrentUser(c).Cart) == 0 {
			return c.Send(cartEmptyMessage, emptyMenu)
		}

		// TODO: save address to db
		getCurrentUser(c).IsSettingAddress = false

		return c.Send(paymentMethodMenuMessage, paymentMethodMenu)
	})

	sendOrder := func(c tele.Context) error {
		// TODO: add order to db

		getCurrentUser(c).Cart = []*model.Product{}

		c.Send(orderMessage)
		return c.Send(menuMessage, menu)
	}

	bot.Handle(&btnCash, func(c tele.Context) error {
		// TODO: set payment method to cash for current user in db

		return sendOrder(c)
	})

	bot.Handle(&btnCreditCard, func(c tele.Context) error {
		// TODO: set payment method to credit card for current user in db

		return sendOrder(c)
	})

	bot.Handle(&btnInfo, func(c tele.Context) error {
		return c.Send(infoMessage, emptyMenu)
	})

	bot.Handle(&btnOrderList, func(c tele.Context) error {
		// TODO: get orders from db

		return c.Send(emptyMessage, emptyMenu)
	})

	bot.Handle(&btnContactManager, func(c tele.Context) error {
		return c.Send(managerMessage, emptyMenu)
	})

	bot.Handle(&btnBack, func(c tele.Context) error {
		return c.Send(menuMessage, menu)
	})

	bot.Handle(&btnInlineBack, func(c tele.Context) error {
		c.Send(menuMessage, menu)
		return c.Respond()
	})
}

func priceSum(products []*model.Product) int {
	var sum int
	for _, product := range products {
		sum += product.Price
	}
	return sum
}
