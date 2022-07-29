package handlers

import (
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/view"
	tele "gopkg.in/telebot.v3"
)

func NewStoreHandler(service adapter.ServiceContainer) *StoreHandler {
	return &StoreHandler{service: service}
}

type StoreHandler struct {
	service adapter.ServiceContainer
}

func (h *StoreHandler) SendInfo(c tele.Context) error {
	return c.Send(view.InfoMessage, view.EmptyMenu)
}
