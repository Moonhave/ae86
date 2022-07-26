package handlers

import (
	"ae86/internal/enums"
	"ae86/internal/model"
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/view"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

type CategoryHandler struct {
	service adapter.ServiceContainer
}

func NewCategoryHandler(service adapter.ServiceContainer) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAllCategories() (categories []model.Category, err error) {
	return h.service.Category().ListAll()
}

func (h *CategoryHandler) SendCategories(c tele.Context) error {
	categories, err := h.service.Category().ListAll()
	if err != nil {
		return err
	}

	categoryMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
	categoryMenuRows := make([]tele.Row, 0)

	for _, category := range categories {
		categoryID := category.ID
		btn := categoryMenu.Text(category.Title)
		categoryMenuRows = append(categoryMenuRows, categoryMenu.Row(btn))
		c.Bot().Handle(&btn, func(ctx tele.Context) error {
			products, err := h.service.Product().ListByCategoryID(categoryID)
			if err != nil {
				return err
			}

			return h.sendProductsByCategoryID(ctx, products)
		})
	}

	categoryMenuRows = append(categoryMenuRows, categoryMenu.Row(view.BtnCategoryBack))

	categoryMenu.Reply(categoryMenuRows...)
	return c.Send(view.CategoryMenuMessage, categoryMenu)
}

func (h *CategoryHandler) sendProductsByCategoryID(c tele.Context, products []model.Product) error {
	var messages []tele.Message

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
		c.Bot().Handle(&btnAddToCart, func(c tele.Context) error {
			return h.handleAddToCartButton(c, products)
		})

		messages = append(messages, tele.Message{
			Text:        text,
			ReplyMarkup: productInfoMenu,
			Photo:       &tele.Photo{File: tele.FromURL(p.Image)},
		})
	}

	for i := range messages {
		c.Send(messages[i].Photo)
		c.Send(messages[i].Text, messages[i].ReplyMarkup)
	}
	return c.Respond()
}

func (h *CategoryHandler) handleAddToCartButton(c tele.Context, products []model.Product) error {
	numMenu := &tele.ReplyMarkup{ResizeKeyboard: true}

	var buttonRows []tele.Row
	var currentRow []tele.Btn

	productIndex, _ := strconv.Atoi(c.Args()[0])
	product := &products[productIndex]

	for i := 1; i <= 6; i++ {
		btn := numMenu.Data(fmt.Sprintf("%d", i), fmt.Sprintf("add_product_%d_%d", product.ID, i), fmt.Sprintf("%d", i))
		currentRow = append(currentRow, btn)
		if len(currentRow) == 3 {
			buttonRows = append(buttonRows, numMenu.Row(currentRow...))
			currentRow = []tele.Btn{}
		}
		orderItem := model.OrderItem{ProductID: product.ID, Amount: i}
		c.Bot().Handle(&btn, func(c tele.Context) error {
			customer, err := h.service.Customer().ByExternalID(uint(c.Sender().ID))
			if err != nil {
				return err
			}

			state := enums.OrderStatePending
			orders, err := h.service.Order().ListBy(adapter.OrderFilter{CustomerID: &customer.ID, State: &state})
			if err != nil {
				return err
			}

			var orderId uint
			if len(orders) == 0 {
				orderId, err = h.service.Order().Create(model.Order{CustomerID: customer.ID, State: state})
				if err != nil {
					return err
				}
			} else {
				orderId = orders[0].ID
			}

			orderItem.OrderID = orderId

			_, err = h.service.OrderItem().Create(orderItem)
			if err != nil {
				return err
			}
			return updateProductInlineMenu(c)
		})
	}
	numMenu.Inline(buttonRows...)
	c.Edit(numMenu)

	return c.Respond(&tele.CallbackResponse{Text: view.SelectAmountMessage})
}

func updateProductInlineMenu(c tele.Context) error {
	buttonRows := []tele.Row{
		view.ProductMenu.Row(view.BtnInlineAdded),
		view.ProductMenu.Row(view.BtnInlineOrder),
		view.ProductMenu.Row(view.BtnInlineBack),
	}
	view.ProductMenu.Inline(buttonRows...)
	c.Edit(view.ProductMenu)

	return c.Respond(&tele.CallbackResponse{Text: view.AddedToCartMessage})
}
