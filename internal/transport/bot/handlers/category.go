package handlers

import (
	"ae86/internal/model"
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/view"
	"gopkg.in/telebot.v3"
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

func (h *CategoryHandler) SendCategories(c telebot.Context) error {
	categories, err := h.service.Category().ListAll()
	if err != nil {
		return err
	}

	categoryMenu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	categoryMenuRows := make([]telebot.Row, 0)

	for _, category := range categories {
		btn := categoryMenu.Text(category.Title)
		categoryMenuRows = append(categoryMenuRows, categoryMenu.Row(btn))
		c.Bot().Handle(&btn, func(ctx telebot.Context) error {
			// todo
			return nil
		})
	}

	categoryMenu.Reply(categoryMenuRows...)
	return c.Send(view.CategoryMenuMessage, categoryMenu)
}
