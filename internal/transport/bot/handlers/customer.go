package handlers

import (
	"ae86/internal/enums"
	"ae86/internal/model"
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/temp"
	"ae86/internal/transport/bot/view"
	tele "gopkg.in/telebot.v3"
)

func NewCustomerHandler(service adapter.ServiceContainer) *CustomerHandler {
	return &CustomerHandler{service: service}
}

type CustomerHandler struct {
	service adapter.ServiceContainer
}

// CustomerInfo TempCustomerInfo - temporary customer info, for storing a customer's cart and other temporary data
type CustomerInfo struct {
	Cart                    []*model.OrderItem
	IsRequiredToSendAddress bool
	PreferredAddress        string
	PreferredPaymentMethod  enums.PaymentMethod
}

func (h *CustomerHandler) CreateCustomer(c tele.Context) error {
	temp.CreateCustomer(c)
	return c.Send(view.MenuMessage, view.Menu)
}

func (h *CustomerHandler) TryStoreAddress(c tele.Context) error {
	if !temp.GetCurrentCustomer(c).IsRequiredToSendAddress {
		return c.Send(view.UnknownCommandMessage, view.EmptyMenu)
	}

	if len(temp.GetCurrentCustomer(c).Cart) == 0 {
		return c.Send(view.CartEmptyMessage, view.EmptyMenu)
	}

	temp.GetCurrentCustomer(c).PreferredAddress = c.Text()

	// TODO: save address to db
	temp.GetCurrentCustomer(c).IsRequiredToSendAddress = false

	return c.Send(view.PaymentMethodMenuMessage, view.PaymentMethodMenu)
}
