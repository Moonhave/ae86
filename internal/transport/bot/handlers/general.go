package handlers

import (
	"ae86/internal/model"
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

func (h *GeneralHandler) Start(c tele.Context) error {
	customerExternalID := c.Sender().ID

	exists, err := h.service.Customer().ExistsByExternalID(uint(customerExternalID))
	if err != nil {
		return err
	}

	if !exists {
		_, err = h.service.Customer().Create(model.Customer{
			ExternalID: uint(customerExternalID),
			Username:   c.Sender().Username,
			FirstName:  c.Sender().FirstName,
			LastName:   c.Sender().LastName,
		})
		if err != nil {
			return err
		}
	}

	return c.Send(view.MenuMessage, view.Menu)
}

func (h *GeneralHandler) GoBackToMenu(c tele.Context) error {
	err := c.Send(view.MenuMessage, view.Menu)
	if err != nil {
		return err
	}
	return c.Respond()
}
