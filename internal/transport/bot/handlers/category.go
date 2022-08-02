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
	return h.service.Category().GetAllCategories()
}

func (h *CategoryHandler) SendCategories(c telebot.Context) error {
	return c.Send(view.CategoryMenuMessage, view.CategoryMenu)
}
