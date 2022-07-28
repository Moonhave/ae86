package handlers

import (
	"ae86/internal/transport/adapter"
	"gopkg.in/telebot.v3"
)

type CategoryHandler struct {
	service adapter.ServiceContainer
}

func NewCategoryHandler(service adapter.ServiceContainer) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) LoadCategories(c telebot.Context) error {
	// load categories from service and convert to buttons
	return nil
}