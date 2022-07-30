package bot

import (
	"ae86/internal/container"
	"ae86/internal/model"
	"ae86/internal/transport/bot/temp"
	"ae86/internal/transport/bot/view"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"strconv"
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

	LoadCategories(bot, handlers)
	InitializeMenuReplies()
	RegisterCommands(bot, handlers)
	RegisterButtonCallbacks(bot, handlers)
	RegisterEvents(bot, handlers)

	bot.Start()
	return nil
}

func LoadCategories(bot *tele.Bot, handlers *container.HandlerContainer) {
	//var categories = []model.Category{{Title: "Пицца 🍕"}, {Title: "Суши 🍣"}, {Title: "Десерты 🍨"}, {Title: "Напитки 🍹"},}
	categories, err := handlers.Category().GetAllCategories()
	if err != nil {
		return
	}

	for _, value := range categories {
		category := value
		btn := view.CategoryMenu.Text(category.Title)
		view.CategoryMenuRows = append(view.CategoryMenuRows, view.CategoryMenu.Row(btn))
		bot.Handle(&btn, func(c tele.Context) error {
			messages := LoadProducts(bot, category.ID, handlers)
			for i := range messages {
				c.Send(messages[i].Photo)
				c.Send(messages[i].Text, messages[i].ReplyMarkup)
			}
			return c.Respond()
		})
	}

	view.CategoryMenuRows = append(view.CategoryMenuRows, view.CategoryMenu.Row(view.BtnCategoryBack))
}

func LoadProducts(bot *tele.Bot, categoryId uint, handlers *container.HandlerContainer) (messages []tele.Message) {
	/*var products = []model.Product{
		{Title: category.Title + " 1", Description: "Описание", Price: 1190},
		{Title: category.Title + " 2", Description: "Описание", Price: 1490},
		{Title: category.Title + " 3", Description: "Описание", Price: 1790},
	}*/
	products, err := handlers.Product().GetAllProductsByCategoryID(categoryId)
	if err != nil {
		return
	}

	for i, product := range products {
		productInfoMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
		btnAddToCart := productInfoMenu.Data(view.BtnInlineAddMessage, fmt.Sprintf("add_product_%d", i), fmt.Sprintf("%d", i))

		buttonRows := []tele.Row{productInfoMenu.Row(btnAddToCart)}
		isLastProduct := i == len(products)-1
		if isLastProduct {
			buttonRows = append(buttonRows, productInfoMenu.Row(view.BtnInlineBack))
		}
		productInfoMenu.Inline(buttonRows...)

		p := product
		text := fmt.Sprintf("%s\n%s\nЦена: %d тенге", p.Title, p.Description, p.Price)
		HandleAddToCartButton(bot, btnAddToCart, products, handlers)

		messages = append(messages, tele.Message{
			Text:        text,
			ReplyMarkup: productInfoMenu,
			Photo:       &tele.Photo{File: tele.FromURL(p.Image)},
		})
	}
	return messages
}

func HandleAddToCartButton(bot *tele.Bot, btn tele.Btn, products []model.Product, handlers *container.HandlerContainer) {
	bot.Handle(&btn, func(c tele.Context) error {
		numMenu := &tele.ReplyMarkup{ResizeKeyboard: true}

		var buttonRows []tele.Row
		var currentRow []tele.Btn

		productIndex, _ := strconv.Atoi(c.Args()[0])
		product := products[productIndex]

		for i := 1; i <= 6; i++ {
			btn := numMenu.Data(fmt.Sprintf("%d", i), fmt.Sprintf("add_product_%d_%d", product.ID, i), fmt.Sprintf("%d", i))
			currentRow = append(currentRow, btn)
			if len(currentRow) == 3 {
				buttonRows = append(buttonRows, numMenu.Row(currentRow...))
				currentRow = []tele.Btn{}
			}
			orderItem := &model.OrderItem{Product: &product, Amount: i}
			bot.Handle(&btn, handlers.Product().UpdateProductInlineMenu, func(handlerFunc tele.HandlerFunc) tele.HandlerFunc {
				temp.AddToCart(c, orderItem)
				return handlerFunc
			})
		}
		numMenu.Inline(buttonRows...)
		c.Edit(numMenu)

		return c.Respond(&tele.CallbackResponse{Text: view.SelectAmountMessage})
	})
}

func InitializeMenuReplies() {
	view.Menu.Reply(
		view.Menu.Row(view.BtnCategories),
		view.Menu.Row(view.BtnCart, view.BtnOrder),
		view.Menu.Row(view.BtnInfo),
		view.Menu.Row(view.BtnOrderList),
		view.Menu.Row(view.BtnContactManager),
	)

	view.CategoryMenu.Reply(view.CategoryMenuRows...)

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
