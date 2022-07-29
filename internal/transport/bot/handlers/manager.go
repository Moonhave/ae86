package handlers

import (
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/view"
	tele "gopkg.in/telebot.v3"
)

func NewManagerHandler(service adapter.ServiceContainer) *ManagerHandler {
	return &ManagerHandler{service: service}
}

type ManagerHandler struct {
	service adapter.ServiceContainer
}

func (h *ManagerHandler) SendManagerDetails(c tele.Context) error {
	return c.Send(view.ManagerMessage, view.EmptyMenu)
}
