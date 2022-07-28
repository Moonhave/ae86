package bot

import (
	"ae86/internal/container"
	"ae86/internal/model"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

// TempUserInfo - temporary user info, for storing a user's cart and other data
type TempUserInfo struct {
	Cart             []*model.OrderItem
	IsSettingAddress bool
}

// temp user storage
var userStorage = make(map[int64]*TempUserInfo)

func getCurrentUser(c tele.Context) *TempUserInfo {
	if userStorage[c.Sender().ID] == nil {
		userStorage[c.Sender().ID] = &TempUserInfo{}
	}
	return userStorage[c.Sender().ID]
}

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

	LoadCategories(bot)
	InitializeMenuReplies()
	RegisterEndpointCallbacks(bot, handlers)
	RegisterButtonCallbacks(bot, handlers)

	bot.Start()
	return nil
}

func LoadCategories(bot *tele.Bot) {
	// TODO: populate with actual categories from db
	var categories = []model.Category{{Title: "Пицца"}, {Title: "Суши"}, {Title: "Десерты"}, {Title: "Напитки"}}

	for _, category := range categories {
		btn := categoryMenu.Text(category.Title)
		categoryMenuRows = append(categoryMenuRows, categoryMenu.Row(btn))

		bot.Handle(&btn, func(c tele.Context) error {
			messages, menus := loadProducts(bot, category)
			for i := range messages {
				c.Send(messages[i], menus[i])
			}
			return c.Respond()
		})
	}

	categoryMenuRows = append(categoryMenuRows, categoryMenu.Row(btnCategoryBack))
}

func loadProducts(bot *tele.Bot, category model.Category) (messages []string, markups []*tele.ReplyMarkup) {
	// TODO: get products from db depending on category

	var testProducts = []model.Product{
		{
			Title:       category.Title + " 1",
			Description: "Описание",
			Price:       1190,
		},
		{
			Title:       category.Title + " 2",
			Description: "Описание",
			Price:       1490,
		},
		{
			Title:       category.Title + " 3",
			Description: "Описание",
			Price:       1790,
		},
	}

	for i, product := range testProducts {
		productInfoMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
		btnAddToCart := productInfoMenu.Data(btnInlineAddMessage, fmt.Sprintf("add_product_%d", i), fmt.Sprintf("%d", i))

		buttonRows := []tele.Row{
			productInfoMenu.Row(btnAddToCart),
		}
		isLastProduct := i == len(testProducts)-1
		if isLastProduct {
			buttonRows = append(buttonRows, productInfoMenu.Row(btnInlineBack))
		}
		productInfoMenu.Inline(buttonRows...)

		p := product
		text := fmt.Sprintf("%s\n%s\nЦена: %d тенге", p.Title, p.Description, p.Price)
		handleAddToCartButton(bot, btnAddToCart, testProducts)

		messages = append(messages, text)
		markups = append(markups, productInfoMenu)
	}
	return messages, markups
}

func handleAddToCartButton(bot *tele.Bot, btn tele.Btn, products []model.Product) {
	numMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
	bot.Handle(&btn, func(c tele.Context) error {
		var buttonRows []tele.Row
		var currentRow []tele.Btn

		productIndex, _ := strconv.Atoi(c.Args()[0])
		product := products[productIndex]

		for i := 1; i <= 6; i++ {
			btn := numMenu.Data(fmt.Sprintf("%d", i), fmt.Sprintf("add_product_%d_%d", product.ID, i))
			currentRow = append(currentRow, btn)
			if len(currentRow) == 3 {
				buttonRows = append(buttonRows, numMenu.Row(currentRow...))
				currentRow = []tele.Btn{}
			}
			handleAddProductButton(bot, btn, &model.OrderItem{
				Product: &product,
				Amount:  i,
			})
		}
		numMenu.Inline(buttonRows...)
		c.Edit(numMenu)

		return c.Respond(&tele.CallbackResponse{Text: selectAmountMessage})
	})
}

func handleAddProductButton(bot *tele.Bot, btn tele.Btn, orderItem *model.OrderItem) {
	bot.Handle(&btn, func(c tele.Context) error {
		getCurrentUser(c).Cart = append(getCurrentUser(c).Cart, orderItem)
		buttonRows := []tele.Row{
			productMenu.Row(btnInlineAdded),
			productMenu.Row(btnInlineOrder),
			productMenu.Row(btnInlineBack),
		}
		productMenu.Inline(buttonRows...)
		c.Edit(productMenu)

		return c.Respond(&tele.CallbackResponse{Text: addedToCartMessage})
	})
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

func RegisterEndpointCallbacks(bot *tele.Bot, handlers *container.HandlerContainer) {
	bot.Handle("/start", func(c tele.Context) error {
		userStorage[c.Sender().ID] = &TempUserInfo{
			Cart:             []*model.OrderItem{},
			IsSettingAddress: false,
		}
		return c.Send(menuMessage, menu)
	})
}

func RegisterButtonCallbacks(bot *tele.Bot, handlers *container.HandlerContainer) {
	bot.Handle(&btnCategories, func(c tele.Context) error {
		return c.Send(categoryMenuMessage, categoryMenu)
	})

	bot.Handle(&btnCart, func(c tele.Context) error {
		if len(getCurrentUser(c).Cart) == 0 {
			return c.Send(cartEmptyMessage, emptyMenu)
		}
		text := ""
		for _, orderItem := range getCurrentUser(c).Cart {
			product := orderItem.Product
			text += fmt.Sprintf("%s\n%s\nЦена: %dx%d=%d тенге\n\n", product.Title, product.Description,
				product.Price, orderItem.Amount, product.Price*orderItem.Amount)
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

		getCurrentUser(c).Cart = []*model.OrderItem{}

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

func priceSum(products []*model.OrderItem) (sum int) {
	for _, product := range products {
		sum += product.Product.Price * product.Amount
	}
	return sum
}
