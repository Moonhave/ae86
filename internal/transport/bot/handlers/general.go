package handlers

import (
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/view"
	tele "gopkg.in/telebot.v3"
)

func NewGeneralHandler(service adapter.ServiceContainer) *GeneralHandler {
	return &GeneralHandler{service: service}
}

type GeneralHandler struct {
	service adapter.ServiceContainer
}

func (h *GeneralHandler) GoBackToMenu(c tele.Context) error {
	return c.Send(view.MenuMessage, view.Menu)
}

func (h *GeneralHandler) GoBackToMenuInline(c tele.Context) error {
	err := h.GoBackToMenu(c)
	if err != nil {
		return err
	}
	return c.Respond()
}
