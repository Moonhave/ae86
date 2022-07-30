package handlers

import (
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/view"
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func NewManagerHandler(service adapter.ServiceContainer) *ManagerHandler {
	return &ManagerHandler{service: service}
}

type ManagerHandler struct {
	service adapter.ServiceContainer
}

func (h *ManagerHandler) SendManagerDetails(c tele.Context) error {
	manager, err := h.service.Manager().GetManager()
	if err != nil {
		c.Send(view.DefaultManagerMessage, view.EmptyMenu)
		return err
	}
	return c.Send(fmt.Sprintf("Контакт менеджера: %s %s\n%s", manager.FirstName, manager.LastName,
		manager.Phone), view.EmptyMenu)
}
